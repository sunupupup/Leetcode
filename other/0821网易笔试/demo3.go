package main

import (
	"fmt"
)

func test3() {
	//先从最小的人发纸

	//记录已经发了纸的人

	children := []int{1, 1, 1}

	all := len(children)
	mask := make([]int, all)

	var goLeft func(int)
	goLeft = func(index int) {
		for index-1 >= 0 {
			if children[index-1] > children[index] {
				mask[index-1] = max(mask[index]+1, mask[index-1])
			} else if children[index-1] == children[index] {
				mask[index-1] = mask[index]
			} else {
				break
			}
			index--
		}
	}
	var goRight func(int)
	goRight = func(index int) {
		for index+1 <= all-1 {
			if children[index+1] > children[index] {
				mask[index+1] = max(mask[index]+1, mask[index+1])
			} else if children[index+1] == children[index] {
				mask[index+1] = mask[index]
			} else {
				break
			}
			index++
		}
	}

	//检查是否还有人没被分配
	for check(&mask) {
		//找到最小的年龄的人的下标
		smallChildren := findMinAge(&children, &mask)
		//fmt.Println(smallChildren)
		//从每个最小孩子向两边遍历
		for i := 0; i < len(smallChildren); i++ {
			mask[smallChildren[i]] = 1
			now := smallChildren[i]
			//向左右遍历
			goLeft(now)
			goRight(now)
		}
		//fmt.Println(mask)
	}
	sum := 0
	for i := 0; i < len(mask); i++ {
		sum += mask[i]
	}
	fmt.Println(sum)
}

func check(mask *[]int) bool {
	for i := 0; i < len(*mask); i++ {
		if (*mask)[i] == 0 {
			return true
		}
	}
	return false
}

func findMinAge(children *[]int, mask *[]int) []int {

	minAge := 100
	for i := 0; i < len(*children); i++ {
		if (*children)[i] < minAge && (*mask)[i] == 0 {
			minAge = (*children)[i]
		}
	}

	ret := []int{}
	for i := 0; i < len(*children); i++ {
		if (*children)[i] == minAge && (*mask)[i] == 0 {
			ret = append(ret, i)
		}
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

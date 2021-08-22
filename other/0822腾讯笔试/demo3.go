package main

import (
	"fmt"
	"sort"
)

func main() {

	//二分
	T := 0
	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		n, w := 0, 0
		fmt.Scan(&n, &w)
		nums := []int{}
		num := 0
		for i := 0; i < n; i++ {
			fmt.Scan(&num)
			nums = append(nums, num)
		}
		sort.Ints(nums)
		//需要小舟的数量 1 - n  用二分法进行计算
		var isValid func(number int) bool
		isValid = func(number int) bool {
			//检查这么多小船能不能满足这些人
			mask := make([]int, n)
			for i := len(nums) - 1; i >= 0; i-- {
				if mask[i] == 0 {
					//找这个人匹配的小船
					if nums[i] == w {
						mask[i] = 1
						number--
						if number < 0 {
							return false
						}
					} else {
						//找之和匹配的
						temp := number
						for j := 0; j < i; j++ {
							if mask[j] != 0 {
								continue
							}
							if (nums[i]+nums[j])%2 == 0 && nums[i]+nums[j] <= w {
								mask[i], mask[j] = 1, 1
								number--
								if number < 0 {
									return false
								}
								break
							} else if nums[i]+nums[j] > w {
								break
							}
						}
						if temp == number {
							mask[i] = 1
							number--
							if number < 0 {
								return false
							}
						}
					}
				}
			}
			return true
		}

		if len(nums)==1 {
			fmt.Println(1)
			continue
		}

		l, r := 1, len(nums)
		for l < r {
			mid := (l + r) / 2
			if isValid(mid) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		fmt.Println(l)
	}

}

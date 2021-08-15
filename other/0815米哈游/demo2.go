package main

import "fmt"

func main() {
	//找到奇数位和偶数为的众数即可
	n := 0
	fmt.Scan(&n)
	nums := make([]int, n, n)
	for i := 0; i < n; i++ {
		temp := 0
		fmt.Scan(&temp)
		nums[i] = temp
	}

	//找偶数位的众数
	ret := 0
	m1 := map[int]int{}

	for i := 1; i < n; i += 2 {
		m1[nums[i]]++
	}

	maxCount := 0
	for _, v := range m1 {
		if maxCount < v {
			maxCount = v
		}
	}

	ret += n/2 - maxCount

	m2 := map[int]int{}
	for i := 0; i < n; i += 2 {
		m2[nums[i]]++
	}
	maxCount = 0
	for _, v := range m2 {
		if maxCount < v {
			maxCount = v
		}
	}
	if n%2 == 1{
		ret += n/2+1 - maxCount
	}else{
		ret += n/2 - maxCount
	}
	fmt.Println(ret)
}

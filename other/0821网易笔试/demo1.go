package main

import (
	"fmt"
)

func test1() {
	temp := []int{}
	num := 0
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		temp = append(temp, num)
	}

	n := len(temp) - 1
	nums := temp[:n]
	target := temp[n]

	ret := 0
	for i := 0; i < n; i++ {
		a := nums[i]
		for j := i + 1; j < n; j++ {
			b := nums[j]
			if a+b <= target {
				ret++
			}
		}
	}

	fmt.Println(ret)
}

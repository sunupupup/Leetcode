package main

import (
	"fmt"
	"sort"
)

func demo2() {

	T := 0
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		n := 0
		fmt.Scan(&n)
		nums := []int{}
		num := 0
		for i := 0; i < n; i++ {
			fmt.Scan(&num)
			nums = append(nums, num)
		}

		ret := 0
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] > nums[j]
		})

		temp := 1
		for i := n-1; i >= 0; i-- {
            ret = (ret + (nums[i] * temp) % 1000000007) % 1000000007
			temp = (temp * 2) % 1000000007
		}
		fmt.Println(ret % 1000000007)
	}

}

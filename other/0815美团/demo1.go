package main

import (
	"fmt"
	"sort"
)

/*
小美给小团一个n个数字构成的数字序列，问小团能不能经过重新排列后形成1到n的排列。

举例：

小美给小团[2, 1, 3]，则可以经过重新排列后构成[1, 2, 3]，这是可行的。

小美给小团[4, 4, 1, 3]，则无法经过重新排列后构成[1, 2, 3, 4]，这是不可行的。

为了防止小团靠运气碰对答案，小美会进行多组询问。



输入描述
第一行是一个数T，表示有T组数据。

对于每组数据：

第一行一个数字n表示小美给出的序列由n个数字构成。

接下来一行n个空格隔开的正整数。

*/
func test1() {

	T := 0
	fmt.Scan(&T)

	for i := 0; i < T; i++ {

		n := 0
		fmt.Scan(&n)
		nums := make([]int, 0, n)
		for a := 0; a < n; a++ {
			temp := 0
			fmt.Scan(&temp)
			nums = append(nums, temp)
		}

		//检查nums
		sort.Ints(nums)
		check(&nums)

	}
}

func check(nums *[]int) {
	for i, v := range *nums {
		if v != i+1 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")

}

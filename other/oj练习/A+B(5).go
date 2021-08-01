/*
输入的第一行包括一个正整数t(1 <= t <= 100), 表示数据组数。
接下来t行, 每行一组数据。
每行的第一个整数为整数的个数n(1 <= n <= 100)。
接下来n个正整数, 即需要求和的每个正整数。

2
4 1 2 3 4
5 1 2 3 4 5
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func test5_1() {
	var t, n, num int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&n)
		res := 0
		for ; n > 0; n-- {
			fmt.Scan(&num)
			res += num
		}
		fmt.Println(res)
	}
}

// 当然这个题可以使用bufio包来做，也就是读入一行数据再进行处理

func test5_2() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	t, _ := strconv.Atoi(input.Text())
	for ; t > 0; t-- {
		input.Scan()
		nums := strings.Split(input.Text(), " ")
		res := 0
		for i := 1; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])
			res += num
		}
		fmt.Println(res)
	}
}

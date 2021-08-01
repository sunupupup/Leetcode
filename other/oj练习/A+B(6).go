/*
输入数据有多组, 每行表示一组输入数据。
每行的第一个整数为整数的个数n(1 <= n <= 100)。
接下来n个正整数, 即需要求和的每个正整数。

4 1 2 3 4
5 1 2 3 4 5
*/

// 这个直接用bufio做比较方便
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func test6_1() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := strings.Split(input.Text(), " ")
		res := 0
		for i := 1; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])
			res += num
		}
		fmt.Println(res)
	}
}

// 当然， fmt。Sacn也可以做， 与第一道题类似的判断就是了
func test6_2() {
	var n, num int
	for {
		if _, err := fmt.Scan(&n); err != io.EOF {
			res := 0
			for ; n > 0; n-- {
				fmt.Scan(&num)
				res += num
			}
			fmt.Println(res)
		} else {
			break
		}
	}
}

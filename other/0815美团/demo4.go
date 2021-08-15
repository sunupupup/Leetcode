package main

import "fmt"

/*
一行一个正整数表示小美最少花费的体力值。

样例输入
10 5
oxxoooxxxo
1 6 9 15 18
*/
func main() {

	sn := 0
	n := 0
	fmt.Scan(&sn)
	fmt.Scan(&n)

	s := ""
	fmt.Scan(&s)
	steps := make([]int, n)
	for i := 0; i < n; i++ {
		temp := 0
		fmt.Scan(&temp)
		steps[i] = temp
	}

	//统计所有水坑的长度
}

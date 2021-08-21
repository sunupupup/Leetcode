package main

import (
	"fmt"
	"strconv"
)

func numDecoding(s string) int {
	bs := []byte(s)
	if len(bs) == 0 || bs[0] == '0' {
		return 1
	}
	//动态规划
	n := len(bs)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {

		if bs[i-1] >= '1' && bs[i-1] <= '9' {
			dp[i] += dp[i-1]
		}
		//判断这两个数是否能结合
		temp := bs[i-2 : i]
		num, _ := strconv.Atoi(string(temp))

		if num >= 10 && num <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
func main() {
	fmt.Println(numDecoding("12"))
}

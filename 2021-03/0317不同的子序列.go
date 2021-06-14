package main

//统计在s中出现多少次t 字串
func numDistinct(s string, t string) int {
	len_s := len(s)
	len_t := len(t)
	if len_t > len_s {
		return 0
	}
	dp := make([][]int, len_t+1)
	for i := range dp {
		dp[i] = make([]int, len_s+1)
	}

	//边界条件
	for i := 0; i <= len_s; i++ {
		dp[len_t][i] = 1
	}

	//动态规划公式
	for i := len_s - 1; i >= 0; i-- {
		for j := len_t - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[j][i] = dp[j+1][i+1] + dp[j][i+1]
			} else {
				dp[j][i] = dp[j][i+1]
			}
		}
	}
	return dp[0][0]
}

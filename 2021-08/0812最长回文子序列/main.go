/*
516. 最长回文子序列
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

示例 1：

输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
*/

func longestPalindromeSubseq(s string) int {
	//if(nums[i]==nums[j])   dp[i][j] = dp[i+1][j-1] + 2
	//if(nums[i]!=nums[j])   dp[i][j] = max(dp[i][j-1],dp[i+1][j])
	n := len(s)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		si := s[i]
		for j := i + 1; j < n; j++ {
			sj := s[j]
			if si == sj {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	//fmt.Println(dp)
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
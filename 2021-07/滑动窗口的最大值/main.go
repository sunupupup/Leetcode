
//dp
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, len(nums)+1)
	}
	for i := 0; i <= len(nums); i++ {
		dp[0][i] = -99999999
	}
	for i := 1; i <= k; i++ {
		for j := 1; j <= len(nums); j++ {
			dp[i][j] = max(dp[i-1][j-1], nums[j-1])
		}
	}

	return dp[k][k:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//最大值，可以考虑优先级队列
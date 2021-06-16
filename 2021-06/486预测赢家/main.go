//dp[i][j]=max(nums[i]−dp[i+1][j],nums[j]−dp[i][j−1])
func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = nums[i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
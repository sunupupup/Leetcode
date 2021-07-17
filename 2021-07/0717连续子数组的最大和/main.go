/*
剑指 Offer 42. 连续子数组的最大和
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。


*/

func maxSubArray(nums []int) int {
	//前缀和  sum[i] 表示前i和数的和
	//子数组和 =
	n := len(nums)
	sum := make([]int, n)
	sum[0] = nums[0]
	minSum := sum[0]
	ret := sum[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + nums[i]
		ret = max(max(ret, sum[i]-minSum), sum[i]) //考虑前面的最小和为负数、正数
		if minSum > sum[i] {
			minSum = sum[i] //记录前面最小的和
		}

	}
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//官方  动态规划
// f(i)=max{f(i−1)+nums[i],nums[i]}
func maxSubArray2(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

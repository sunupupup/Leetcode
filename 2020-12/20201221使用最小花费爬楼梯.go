package main

func main() {

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n < 2 {
		return 0
	}
	var dp0 = 0
	var dp1 = 0
	for i := 2; i < n; i++ {
		temp := dp1
		dp1 = min(dp1+cost[i-1], dp0+cost[i-2])
		dp0 = temp
	}
	return min(dp0+cost[n-2], dp1+cost[n-1])
}

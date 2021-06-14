package test2021

func max(a int,b int)(ret int){
    if a>b {
        ret = a
    }else{
        ret = b
    }
    return
}
func maxProfit(prices []int) int {
	n := len(prices)
	//初始化二维数组
    var dp [5][]int
    for i:=0;i<5;i++{
        dp[i] = make([]int,n,n)
    }
    dp[0][0] = 0
    dp[1][0] = -prices[0]
    dp[2][0] = 0
    dp[3][0] = -prices[0]
    dp[4][0] = 0
    for i:= 1;i<len(prices);i++ {
        dp[0][i] = dp[0][i-1]
        dp[1][i] = max(dp[0][i-1]-prices[i],dp[1][i-1])
        dp[2][i] = max(dp[1][i-1]+prices[i],dp[2][i-1])
        dp[3][i] = max(dp[2][i-1]-prices[i],dp[3][i-1])
        dp[4][i] = max(dp[3][i-1]+prices[i],dp[4][i-1])
    }
    return dp[4][len(prices)-1]
}
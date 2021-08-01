/*


 */

//一开始的做法，深搜   超时了，思考dp，一般深搜会有很多重复的东西
//从主数组中，从前往后一次找再target数组中的数
func minOperations(target []int, arr []int) int {
	//找到最长的  公共子序列  长度为m
	//target的长度为n   由题意target不含重复元素
	//结果为 n - m

	//深搜  还是  dp
	m := make(map[int]int, 0)
	for i, v := range target {
		m[v] = i
	}
	n := len(target)

	ret := 0
	//发现下一个目标数
	//从now开始往后搜，找到下一个出现在target中的字符，且字符出现的下标大于preIndex
	var dfs func(int, int, int)
	dfs = func(preIndex int, now int, count int) {
		length := len(arr)
		i := now + 1
		for ; i < length; i++ {
			if v, ok := m[arr[i]]; ok {
				if v > preIndex { //后面出现的数的index一定大于上次出先的数的index，这个index指的是某个字符在target中的位置
					//dfs(v, i, count+1)
					//dfs(preIndex, i, count)
					//一点点优化，还是超时
					if v == preIndex+1 || v == 1 {
						dfs(v, i, count+1)
					} else {
						dfs(v, i, count+1)
						dfs(preIndex, i, count)
					}
					return
				}
			}
		}
		if i >= length-1 {
			ret = max(ret, count)
			return
			//搜到尾巴了
		}
	}

	//要记录一个preIndex
	dfs(-1, -1, 0)

	return n - ret

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//dp
func minOperations(target []int, arr []int) int {
	n := len(target)
	m := len(arr)
	//dp[m][n]

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		dp[i][m] = 1
	}
	for i := 0; i <= m; i++ {
		dp[n][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if arr[j] == target[i] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return n - (dp[0][0] - 1)
}

//简化的dp，只需要m的复杂度
//但这是普通的最长公共子串的做法，时间复杂度是 O(nm) 本题的n和m长达10^5
//所以需要  二分，对于这么长的数组，得二分解决
func minOperations(target []int, arr []int) int {
	n := len(target)
	m := len(arr)
	//dp[m][n]

	dp := make([]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		temp := make([]int, m+1)
		temp[m] = 1
		for j := m - 1; j >= 0; j-- {
			if arr[j] == target[i] {
				temp[j] = dp[j+1] + 1
			} else {
				temp[j] = max(dp[j], temp[j+1])
			}
		}
		dp = temp
	}

	return n - (dp[0] - 1)

}

//官方：贪心 + 二分
//利用下标代替target数组，转化为arr数组的最长递增子序列问题
func minOperations(target, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n)
	for i, val := range target {
		pos[val] = i
	}
	d := []int{}
	for _, val := range arr {
		if idx, has := pos[val]; has {
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
			} else {
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}

/*
1713. 得到子序列的最少操作次数
给你一个数组 target ，包含若干 互不相同 的整数，以及另一个整数数组 arr ，arr 可能 包含重复元素。

每一次操作中，你可以在 arr 的任意位置插入任一整数。比方说，如果 arr = [1,4,1,2] ，那么你可以在中间添加 3 得到 [1,4,3,1,2] 。
你可以在数组最开始或最后面添加整数。

请你返回 最少 操作次数，使得 target 成为 arr 的一个子序列。

一个数组的 子序列 指的是删除原数组的某些元素（可能一个元素都不删除），同时不改变其余元素的相对顺序得到的数组。
比方说，[2,7,4] 是 [4,2,3,7,2,1,4] 的子序列（加粗元素），但 [2,4,2] 不是子序列。


转化为最长上升子序列问题，见300. 最长递增子序列
*/
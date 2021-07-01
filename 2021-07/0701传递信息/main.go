/*
小朋友 A 在和 ta 的小伙伴们玩传信息游戏，游戏规则如下：

有 n 名玩家，所有玩家编号分别为 0 ～ n-1，其中小朋友 A 的编号为 0
每个玩家都有固定的若干个可传信息的其他玩家（也可能没有）。传信息的关系是单向的（比如 A 可以向 B 传信息，但 B 不能向 A 传信息）。
每轮信息必须需要传递给另一个人，且信息可重复经过同一个人
给定总玩家数 n，以及按 [玩家编号,对应可传递玩家编号] 关系组成的二维数组 relation。返回信息从小 A (编号 0 ) 经过 k 轮传递到编号为 n-1 的小伙伴处的方案数；若不能到达，返回 0。

输入：n = 5, relation = [[0,2],[2,1],[3,4],[2,3],[1,4],[2,0],[0,4]], k = 3

输出：3

解释：信息从小 A 编号 0 处开始，经 3 轮传递，到达编号 4。共有 3 种方案，分别是 0->2->0->4， 0->2->1->4， 0->2->3->4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/chuan-di-xin-xi
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package Solution0701

var ret = 0

func numWays(n int, relation [][]int, k int) int {
	ret = 0
	temp := make([][]int, n) //这个temp就是构造 某个数对应的所有路径，然后深搜即可
	a, b := 0, 0
	for _, v := range relation {
		a = v[0]
		b = v[1]
		temp[a] = append(temp[a], b)
	}

	//深搜
	//从0开始搜
	dfs(temp, 0, 0, n-1, k)
	return ret
}

func dfs(temp [][]int, now int, step int, target int, k int) {
	if step > k {
		return
	}
	step++
	for _, v := range temp[now] {
		if v == target {
			if step == k {
				ret++
				return
			}
		}
		dfs(temp, v, step, target, k)
	}
}

//官方做法：用闭包
func numWays1(n int, relation [][]int, k int) int {
	temp := make([][]int, n)
	a, b := 0, 0
	for _, v := range relation {
		a = v[0]
		b = v[1]
		temp[a] = append(temp[a], b)
	}
	//深搜
	//从0开始搜
	ret := 0
	var dfs func(int, int)
	dfs = func(now, step int) {
		if step > k {
			return
		}
		step++
		for _, v := range temp[now] {
			if v == n-1 {
				if step == k {
					ret++
					return
				}
			}
			dfs(v, step)
		}
	}
	dfs(0, 0)
	return ret
}

//广搜，queue的思维，C++中用queue，Go中直接用切片即可
func numWaysBFS(n int, relation [][]int, k int) int {
	temp := make([][]int, n)
	a, b := 0, 0
	for _, v := range relation {
		a = v[0]
		b = v[1]
		temp[a] = append(temp[a], b)
	}
	//宽搜，搜到第k层时，检查有多少个n-1	每次遍历一层，得到下一层
	var q = []int{0}
	for i := 0; i < k; i++ {
		if len(q) > 0 {
			var qq []int
			for _, v := range q {
				qq = append(qq, temp[v]...)
			}
			q = qq
		}
	}
	ret := 0
	for _, v := range q {
		if v == n-1 {
			ret++
		}
	}
	return ret
}

//因为dfs这种大多存在重复的情况，所以可以多考虑dp
//动态规划	dp[i][j] 表示经过i轮到j位置的走法
func numWaysDP(n int, relation [][]int, k int) int {
	//dp
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < k; i++ {
		for _, r := range relation {
			src, dst := r[0], r[1]
			dp[i+1][dst] += dp[i][src] //dp公式
		}
	}

	return dp[k][n-1]
}

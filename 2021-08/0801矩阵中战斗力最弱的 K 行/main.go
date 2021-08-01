/*
1337. 矩阵中战斗力最弱的 K 行
给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。

请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。

如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。

军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。

示例 1：

输入：mat =
[[1,1,0,0,0],
 [1,1,1,1,0],
 [1,0,0,0,0],
 [1,1,0,0,0],
 [1,1,1,1,1]],
k = 3
输出：[2,0,3]
解释：
每行中的军人数目：
行 0 -> 2
行 1 -> 4
行 2 -> 1
行 3 -> 2
行 4 -> 5
从最弱到最强对这些行排序后得到 [2,0,3,1,4]
*/
package main

import (
	"container/heap"
	"sort"
)

//二分计算每行的战力
func kWeakestRows(mat [][]int, k int) []int {

	//m行 n列
	m := len(mat)
	n := len(mat[0])
	power := map[int][]int{}

	//sort.Search(len(row), func(j int) bool { return row[j] == 0 })   可以直接用这个，标准库的二分
	var f func(int) int
	f = func(i int) int {
		if mat[i][n-1] == 1 {
			return n
		}
		left, right := 0, n-1
		for left < right {
			mid := (left + right) / 2
			if mat[i][mid] == 1 {
				left = mid + 1
			} else {
				right = mid
			}
		}

		return left
	}
	//统计每行的战力
	for i := 0; i < m; i++ {
		power[f(i)] = append(power[f(i)], i)
	}

	ret := []int{}
	for i := 0; i <= n; i++ {
		if v, ok := power[i]; ok {
			ret = append(ret, v...)
		}
	}

	return ret[:k]
}

//官方：二分 + 堆
//当我们得到每一行的战斗力后，我们可以将它们全部放入一个小根堆中，并不断地取出堆顶的元素 k 次，这样我们就得到了最弱的 k 行的索引。
func kWeakestRows2(mat [][]int, k int) []int {
	h := hp{}
	for i, row := range mat {
		pow := sort.Search(len(row), func(j int) bool { return row[j] == 0 })
		h = append(h, pair{pow, i})
	}
	heap.Init(&h) //交由heap.Init()去构建堆
	ans := make([]int, k)
	for i := range ans {
		ans[i] = heap.Pop(&h).(pair).idx
	}
	return ans
}

type pair struct{ pow, idx int }
type hp []pair

//实现接口的方法  堆的len、less等方法都交由用户来实现
func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.pow < b.pow || a.pow == b.pow && a.idx < b.idx
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {

}

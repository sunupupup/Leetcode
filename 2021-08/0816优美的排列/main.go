/*
假设有从 1 到 N 的 N 个整数，如果从这 N 个数字中成功构造出一个数组，使得数组的第 i 位 (1 <= i <= N) 满足如下两个条件中的一个，我们就称这个数组为一个优美的排列。条件：

第 i 位的数字能被 i 整除
i 能被第 i 位上的数字整除
现在给定一个整数 N，请问可以构造多少个优美的排列？

示例1:

输入: 2
输出: 2
解释:

第 1 个优美的排列是 [1, 2]:
  第 1 个位置（i=1）上的数字是1，1能被 i（i=1）整除
  第 2 个位置（i=2）上的数字是2，2能被 i（i=2）整除

第 2 个优美的排列是 [2, 1]:
  第 1 个位置（i=1）上的数字是2，2能被 i（i=1）整除
  第 2 个位置（i=2）上的数字是1，i（i=2）能被 1 整除
说明:

N 是一个正整数，并且不会超过15。	//这里n最大15，所以可以考虑回溯

*/

package main

import "math/bits"

//回溯复杂度是  O(n!)
func countArrangement(n int) int {

	ret := 0
	mask := make([]int, n+1)
	var backtrack func(int)
	backtrack = func(now int) {
		if now > n {
			ret++
			return
		}

		//检查能放那些东西到now这个地方
		for i := 1; i <= n; i++ {
			if mask[i] == 0 && (i%now == 0 || now%i == 0) {
				mask[i] = 1
				backtrack(now + 1)
				mask[i] = 0
			}
		}

	}

	backtrack(1)

	return ret
}

//状态压缩 + 动态规划
func countArrangement2(n int) int {
	f := make([]int, 1<<n)
	f[0] = 1
	for mask := 1; mask < 1<<n; mask++ {
		num := bits.OnesCount(uint(mask))
		for i := 0; i < n; i++ {
			if mask>>i&1 > 0 && (num%(i+1) == 0 || (i+1)%num == 0) {
				f[mask] += f[mask^1<<i]
			}
		}
	}
	return f[1<<n-1]
}

package main

import "sort"

//最简单的，排序，一直取最小的就行了，但是需要sort的复杂度 + 贪心（一直取最小的）
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	count := 0
	for _, v := range costs {
		if coins >= v {
			count++
		}
		coins -= v
	}

	return count
}

//法二：数组计数  不用排序，但内存消耗大很多
/*
执行用时：156 ms, 在所有 Go 提交中击败了97.60%的用户
内存消耗：9.6 MB, 在所有 Go 提交中击败了7.21%的用户
*/
func maxIceCream2(costs []int, coins int) int {
	counts := make([]int, 100001)
	for _, v := range costs {
		counts[v]++
	}

	ret := 0
	for i, v := range counts {
		for v > 0 {
			if coins >= i {
				coins -= i
				v--
				ret++
			} else {
				break
			}
		}
	}

	return ret
}

/*
1833. 雪糕的最大数量
夏日炎炎，小男孩 Tony 想买一些雪糕消消暑。

商店中新到 n 支雪糕，用长度为 n 的数组 costs 表示雪糕的定价，其中 costs[i] 表示第 i 支雪糕的现金价格。Tony 一共有 coins 现金可以用于消费，他想要买尽可能多的雪糕。

给你价格数组 costs 和现金量 coins ，请你计算并返回 Tony 用 coins 现金能够买到的雪糕的 最大数量 。

注意：Tony 可以按任意顺序购买雪糕。



示例 1：

输入：costs = [1,3,2,4,1], coins = 7
输出：4
解释：Tony 可以买下标为 0、1、2、4 的雪糕，总价为 1 + 3 + 2 + 1 = 7

*/

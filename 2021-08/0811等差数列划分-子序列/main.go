/*
446. 等差数列划分 II - 子序列
给你一个整数数组 nums ，返回 nums 中所有 等差子序列 的数目。

如果一个序列中 至少有三个元素 ，并且任意两个相邻元素之差相同，则称该序列为等差序列。

例如，[1, 3, 5, 7, 9]、[7, 7, 7, 7] 和 [3, -1, -5, -9] 都是等差序列。
再例如，[1, 1, 2, 5, 7] 不是等差序列。
数组中的子序列是从数组中删除一些元素（也可能不删除）得到的一个序列。

例如，[2,5,10] 是 [1,2,1,2,4,1,5,10] 的一个子序列。
题目数据保证答案是一个 32-bit 整数。
*/

//暴力法：找出所有的子序列，每个判断下是不是等差
//只能通过 42 / 101 个通过测试用例
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	temp := []int{}
	ret := 0
	var dfs func(int)
	dfs = func(index int) {
		if index == n {
			if helper(&temp) {
				ret++
			}
			return
		}

		dfs(index + 1)
		temp = append(temp, nums[index])
		dfs(index + 1)
		temp = temp[:len(temp)-1]
	}

	dfs(0)

	return ret
}

//判断一个子数组是不是等差
func helper(nums *[]int) bool {
	if len(*nums) < 3 {
		return false
	}

	interval := (*nums)[1] - (*nums)[0]
	for i := 2; i < len(*nums); i++ {
		if (*nums)[i]-(*nums)[i-1] != interval {
			return false
		}
	}
	return true
}

//子序列 ： 动态规划 + 哈希
//首先考虑至少有两个元素的等差子序列，下文将其称作弱等差子序列。
//由于尾项和公差可以确定一个等差数列，定义状态 f[i][d]表示尾项为nums[i]，公差为d的弱等差子序列个数
//用一个二重循环去遍历nums中所有的元素对(nums[i],nums[j],i<j)，将nums[i]和nums[j]看成是等差子序列的后两项，多以公差为 nums[j]-nums[i]
//由此得到了状态转移方程  f[i][d] += f[j][d]
//由于(nums[i],nums[j])这一对也可以看成是弱等差子序列，所以状态转移方程为 f[i][d] += f[j][d]+1
//由于题目要统计的等差子序列至少有三个元素，我们回顾上述二重循环，
//其中「将nums[i] 加到以 nums[j] 为尾项，公差为 d 的弱等差子序列的末尾」这一操作，实际上就构成了一个至少有三个元素的等差子序列，
//因此我们将循环中的 f[j][d] 累加，即为答案。
//代码实现时，由于 nums[i] 的范围很大，所以计算出的公差的范围也很大，我们可以将状态转移数组 f 的第二维用哈希表代替。
func numberOfArithmeticSlices(nums []int) int {
	ret := 0
	f := make([]map[int]int, len(nums))
	for i, x := range nums {
		f[i] = map[int]int{}
		for j, y := range nums[:i] {
			d := x - y
			cnt := f[j][d] //计算以j下标结尾有多少公差为d的子序列
			ret += cnt
			f[i][d] += cnt + 1
		}
	}
	return ret
}

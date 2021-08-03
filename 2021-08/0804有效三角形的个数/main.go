/*
给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。

示例 1:

输入: [2,2,3,4]
输出: 3
解释:
有效的组合是:
2,3,4 (使用第一个 2)
2,3,4 (使用第二个 2)
2,2,3

*/

//排序 + 三层循环   复杂度  n^3
func triangleNumber(nums []int) int {
	//找所有的三个数的排列组合

	//优化：先进行排序,二分找到 最大最小 的数
	sort.Ints(nums)
	ret := 0
	//三层循环
	n := len(nums)
	if n <= 2 {
		return ret
	}

	for i := 0; i < n-2; i++ {
		a := nums[i]
		for j := i + 1; j < n-1; j++ {
			b := nums[j]
			for k := j + 1; k < n; k++ {
				c := nums[k]
				if helper(a, b, c) {
					ret++
				} else {
					break
				}
			}
		}
	}
	return ret
}

func helper(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

//二分  复杂度  n^2*lg(n)   空间复杂度：O(log n))，即为排序需要的栈空间
func triangleNumber(nums []int) int {
	//找所有的三个数的排列组合

	//优化：先进行排序,二分找到 最大最小 的数
	sort.Ints(nums)
	ret := 0
	//三层循环
	n := len(nums)
	if n <= 2 {
		return ret
	}

	for i := 0; i < n-2; i++ {
		a := nums[i]
		for j := i + 1; j < n-1; j++ {
			b := nums[j]
			//去 [j+1 , n-1] 中间找  第一个大于等于 A+B 的索引
			l, r := j+1, n-1
			if nums[n-1] < a+b {
				ret += n - 1 - (j + 1) + 1
			} else {
				//二分找
				for l < r {
					mid := (l + r) / 2
					if nums[mid] < a+b {
						l = mid + 1
					} else {
						r = mid
					}
				}
				ret += l - (j + 1)
			}
		}
	}
	return ret
}

func helper(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

//官方   排序 + 二分
func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			ans += sort.SearchInts(nums[j+1:], v+nums[j])
		}
	}
	return
}

//官方： 排序 + 双指针    时间复杂度   n*lgn + n*(2*n) = n^2
//因为 j 和 k 肯定是递增的
//我们每移动一次 j ，就把k往后移动或者原地不动    j和k保证一直往后移，复杂度降低了
/*
我们使用一重循环枚举 i。当 i 固定时，我们使用双指针同时维护 j 和 k，它们的初始值均为 i；
我们每一次将 j 向右移动一个位置，即 j←j+1，并尝试不断向右移动 k
使得 k 是最大的满足 nums[k]<nums[i]+nums[j] 的下标。我们将max(k−j,0) 累加入答案。
*/
func triangleNumber(nums []int) (ans int) {
	n := len(nums)
	sort.Ints(nums)
	for i, v := range nums {
		k := i
		for j := i + 1; j < n; j++ {
			for k+1 < n && nums[k+1] < v+nums[j] {
				k++
			}
			ans += max(k-j, 0)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

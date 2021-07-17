/*
剑指 Offer 53 - I. 在排序数组中查找数字 I
统计一个数字在排序数组中出现的次数。


示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
*/

func search(nums []int, target int) int {
	//找到第一个数，和第一个大的数
	index1, index2 := 0, 0
	n := len(nums)
	if n == 0 {
		return 0
	}
	left, right := 0, n-1

	if nums[0] > target || target > nums[n-1] {
		return 0
	}

	for right > left {
		mid := (right + left) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	index1 = left
	if nums[index1] != target {
		return 0
	}

	if nums[n-1] == target {
		index2 = n - 1
	} else {
		left, right = 0, n-1
		for right > left {
			mid := (right + left) / 2
			if nums[mid] > target {
				right = mid
			} else {
				left = mid + 1
			}
		}
		index2 = left - 1
	}

	return index2 - index1 + 1
}
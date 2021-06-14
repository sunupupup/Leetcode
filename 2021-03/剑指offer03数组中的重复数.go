package main

import "sort"

//这道题在原书上绝对不是简单级别啊！
//它考察的是程序员的沟通能力，先问面试官要时间/空间需求！！！
//只是时间优先就用字典，
//还有空间要求，就用指针+原地排序数组，
//如果面试官要求空间O(1)并且不能修改原数组，还得写成二分法！！

// 直接用数组记录 记录过的数 就说明出现了重复
func findRepeatNumber(nums []int) int {
	n := len(nums)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		number := nums[i]
		if temp[number] == 0 { //没出现过，记录下来
			temp[number] = 1
		} else { //出现过了，就是答案
			return number
		}
	}
	return -1
}

//方法2 排序 + 查重
func findRepeatNumber1(nums []int) int {
	//特殊判断
	if nums == nil || len(nums) == 0 {
		return -1 //不符合，返回-1
	}
	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return -1
		}
	}

	sort.Ints(nums) // 实现是快排
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			// 前后元素相等
			return nums[i]
		}
	}
	return -1
}

//法3 哈希表   map实现
func findRepeatNumber2(nums []int) int {
	//特殊判断
	if nums == nil || len(nums) == 0 {
		return -1
	}
	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return -1
		}
	}

	m := make(map[int]int, 0)
	for i, v := range nums {
		if _, exist := m[v]; exist { //m[v] 可以直接判断存不存在，不存在就放进去，存在直接返回结果
			// 存在该数
			return v
		}
		// 不存在
		m[v] = i
	}
	return -1
}

//方法4 ：原地置换
/*
三、需要注意到数组中的数字都在 0～n-1 的范围内。如果这个数组中没有重复的数字，那么数组排序之后数字 v 将出现在下标为 i 的位置。
由于数组中有重复的数字，有些位置可能存在多个数字，同时有些位置可能没有数字。

那么我们重排这个数组，从头到尾依次扫描数组中的每个数字。扫描到下标为 i 的数字时，我们首先比较这个数（用 v 表示）是不是等与 i 。
若果是，则接着扫描下一个数字；
如果不是，则拿它和第 v 个数字进行比较。
若果它和第 v 个数字相等，就找到了一个重复的数字（该数字在下标为 i 和 v 的位置都出现了）；
如果它和第 v 个数不相等，就把第 i 个数字和第 v 个数字交换，把 v 放到属于它的位置。
接下来再重复这个比较、交换的过程，直到找到一个重复的数字。

这里以数组 {2, 3, 1, 0, 2, 5, 3} 为例来分析找到重复数字的步骤。
数组的第 0 个数字（从 0 开始计数，和数组的下标保持一致）是 2 ，与它的下标不想等，于是把它和下标为 2 的数字 1 交换，
交换之后的数组是 {1, 3, 2, 0, 2, 5, 3} 。
此时第 0 个数字是 1 ，仍然与它的下标不相等，同样把它和下标为 1 的数字 3 交换，
交换之后的数组是 {3, 1, 2, 0, 2, 5, 3} 。
接下来继续交换第 0 个数字 3 和第 3 个数字 0，
交换之后的数组是 {0, 1, 2, 3, 2, 5, 3} 。
此时第 0 个数字是 0 ，接着扫描下一个数字。在接下来的几个数中，下标为 1，2，3 的数分别为 1，2，3，它们的下标和数字相等，因此继续扫描其后的数字。
接下来扫描下标为 4 的数字 2，由于它的数字和下标不相等，再比较它和下标为 2 的数字，注意到此时数组中下标为 2 的数也是 2 ，也就是数字 2 在下标 2 和 下标 4 的两个位置都出现了，因此找到了一个重复的数字。

*/
//时间复杂度：O(n）
//空间复杂度：O(1)
func findRepeatNumber3(nums []int) int {
	//特殊判断
	if nums == nil || len(nums) == 0 {
		return -1
	}
	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return -1
		}
	}

	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			// swap nums[i] and nums[nums[i]]
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

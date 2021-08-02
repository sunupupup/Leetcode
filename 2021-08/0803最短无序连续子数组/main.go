/*
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。

示例 1：

输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
示例 2：

输入：nums = [1,2,3,4]
输出：0
示例 3：

输入：nums = [1]
输出：0

提示：
1 <= nums.length <= 10^4
-10^5 <= nums[i] <= 10^5
*/

func findUnsortedSubarray(nums []int) int {
	//找头升序到的位置、尾巴降序到的位置
	//记录中间的最大最小值 mid_min  mid_max 、记录头的最大值 head_max、尾的最小值  tail_min
	//双指针从向两边遍历，扩大中间的子数组长度

	n := len(nums)
	head, tail := 0, n-1
	for head < n-1 {
		if nums[head+1] >= nums[head] {
			head++
		} else {
			break
		}
	}

	if head == n-1 {
		return 0
	}

	for tail > 0 {
		if nums[tail-1] <= nums[tail] {
			tail--
		} else {
			break
		}
	}
	fmt.Println(head, tail)

	head_max := nums[head]
	tail_min := nums[tail]
	mid_min, mid_max := 100000+1, -100000-1
	for i := head + 1; i <= tail-1; i++ {
		if nums[i] > mid_max {
			mid_max = nums[i]
		}
		if nums[i] < mid_min {
			mid_min = nums[i]
		}
	}
	if mid_min == 100000+1 {
		mid_min = (head_max + tail_min) / 2
		mid_max = (head_max + tail_min) / 2
	}

	for head_max > mid_min || mid_max > tail_min {

		for head_max > mid_min && head >= 0 {
			//移动头的位置
			if head_max > mid_max {
				mid_max = head_max
			}
			head--
			if head >= 0 {
				head_max = nums[head]
			} else {
				head_max = -10001
			}
		}

		for tail_min < mid_max && tail <= n-1 {

			//移动尾巴的位置
			if tail_min < mid_min {
				mid_min = tail_min
			}
			tail++
			if tail <= n-1 {
				tail_min = nums[tail]
			} else {
				tail_min = 10001
			}
		}
	}

	//双指针形成多个条件成立
	//head_max <= mid_min  &&  mid_max <= tail_min
	return tail - head - 1

}

//最简单的方法，排序，查看与原来一样的位置在哪
func findUnsortedSubarray(nums []int) int {
	numsSorted := append([]int(nil), nums...)
	sort.Ints(numsSorted)
	n, m := -1, len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == numsSorted[i] {
			n = i
		} else {
			break
		}
	}
	if n == len(nums)-1 {
		return 0
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == numsSorted[i] {
			m = i
		} else {
			break
		}
	}
	return m - n - 1

}

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	numsSorted := append([]int(nil), nums...)
	sort.Ints(numsSorted)
	left, right := 0, len(nums)-1
	for nums[left] == numsSorted[left] {
		left++
	}
	for nums[right] == numsSorted[right] {
		right--
	}
	return right - left + 1
}

//官方的一次遍历
//一边遍历，一边更新最大最小值
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	minn, maxn := math.MaxInt64, math.MinInt64
	left, right := -1, -1
	for i, num := range nums {
		if maxn > num {
			right = i
		} else {
			maxn = num
		}
		if minn < nums[n-i-1] {
			left = n - i - 1
		} else {
			minn = nums[n-i-1]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}
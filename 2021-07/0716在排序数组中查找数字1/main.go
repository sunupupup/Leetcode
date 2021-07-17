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
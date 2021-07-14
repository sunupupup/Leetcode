//暴力法 或者 二分法
//也考验对题目的理解和转变

//二分，找到与每个nums[i] 对应的最接近的数
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	//二分法
	n := len(nums1)
	temp := make([]int, 0)
	temp = append(temp, nums1...)
	sort.Ints(temp)
	sum := 0
	for i := 0; i < n; i++ {
		sum += abs(nums1[i] - nums2[i])
	}
	if sum == 0 {
		return 0
	}
	ret := 2 << 31
	for i := 0; i < n; i++ {
		//找到与nums2[i]最接近的nums1[j] 即 temp[index]
		//二分找temp切片
		diff := abs(nums1[i] - nums2[i])
		diff2 := diff
		index := search(temp, nums2[i])
		if index < n && index > 0 {
			diff2 = min(diff2, abs(temp[index]-nums2[i]))
			diff2 = min(diff2, abs(temp[index-1]-nums2[i]))
		} else if index == n {
			diff2 = min(diff2, abs(temp[index-1]-nums2[i]))
		} else if index == 0 {
			diff2 = min(diff2, abs(temp[index]-nums2[i]))
		}
		if diff2 < diff {
			ret = min(ret, sum-(diff-diff2))
		}
	}

	return ret % 1000000007
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//二分找第一个大的数
func search(nums []int, num int) int {

	if num > nums[len(nums)-1] {
		return len(nums)
	}

	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= num {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

//暴力法超时
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	// 一开始的ret = sum(  |num1[i]- num2[i] |)
	//现在要将其中一个进行替换
	// 让其中一个   |num1[i]- num2[i] | 替换成  |num1[j]- num2[i] |
	//题目就变为了 ，找到 最大的 （|num1[j]- num2[i]| - |num1[i]- num2[i]| ） 时 i和j 的值
	jj, ii := 0, 0
	maxdiff := 0
	n := len(nums1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//nums1[j] 替换 nums1[i] 后产生的差 ,找到最大的差
			temp := abs(nums1[i]-nums2[i]) - abs(nums1[j]-nums2[i])
			if temp >= maxdiff {
				maxdiff = temp
				ii = i
				jj = j
			}
		}
	}

	nums1[ii] = nums1[jj]

	ret := 0
	for i := 0; i < n; i++ {
		ret += abs(nums1[i] - nums2[i])
	}
	return ret % 1000000007
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
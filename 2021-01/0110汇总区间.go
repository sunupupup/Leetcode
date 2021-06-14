package test2021

import "strconv"

func summaryRanges1(nums []int) []string {
	var ret []string
	n := len(nums)
	for i := 0; i < n; i++ {
		j := i
		for i < n-1 && nums[i+1] == nums[i]+1 {
			i++
		}
		if i == j {
			ret = append(ret, strconv.Itoa(nums[i]))
		} else {
			ret = append(ret, strconv.Itoa(nums[j])+"->"+strconv.Itoa(nums[i]))
		}

	}
	return ret
}

func summaryRanges2(nums []int) []string {
	var ret []string
	n := len(nums)
	for i := 0; i < n; i++ {
		j := i
		for i < n-1 && nums[i+1] == nums[i]+1 {
			i++
		}

		str := strconv.Itoa(nums[j])

		if i > j {
			str += "->" + strconv.Itoa(nums[i])
		}

		ret = append(ret, str)

	}
	return ret
}

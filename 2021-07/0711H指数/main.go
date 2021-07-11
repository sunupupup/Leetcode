//二分
//找到符合要求的最大一个的位置    等价于  找到不合要求的第一个位置减一
func hIndex(citations []int) int {
	//显然二分  因子的范围在 0- max之间，找到最大的h
	max := 0
	for i := range citations {
		if citations[i] > max {
			max = citations[i]
		}
	}

	left := 0
	right := max + 1 //注意这里最好加个1

	for left < right {
		mid := (right-left)/2 + left
		if helper(citations, mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

func helper(citations []int, h int) (ret bool) {
	count := 0
	for i := range citations {
		if citations[i] >= h {
			count++
		}
	}
	if count >= h {
		ret = true
	} else {
		ret = false
	}
	return
}

//不二分当然也可以
func hIndex(citations []int) int {
	max := 0
	ret := 0
	for i := range citations {
		if citations[i] > max {
			max = citations[i]
		}
	}
	f := func(h int) (count int) {
		for i := range citations {
			if citations[i] >= h {
				count++
			}
		}
		return
	}

	for i := 0; i <= max; i++ {
		if f(i) >= i {
			ret = i
		} else {
			break
		}
	}
	return ret
}
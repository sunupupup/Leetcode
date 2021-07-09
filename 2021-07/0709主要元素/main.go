//哈希表
func majorityElement(nums []int) int {
	m := make(map[int]int, 0)
	for i := range nums {
		m[nums[i]]++
	}
	ret := 0
	n := len(nums)
	for i, v := range m {
		if v > ret {
			ret = v
			if ret >= n/2+1 {
				return i
			}
		}
	}
	return -1
}

//摩尔投票法   数量相抵，个数最多的数肯定活到最后，诅咒检查一下就行
func majorityElement(nums []int) int {
	now := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			now = nums[i]
			count = 1
		} else {
			if nums[i] == now {
				count++
			} else {
				count--
			}
		}
	}
	if count == 0 {
		return -1
	}

	//检查now这个数字个数是否超过一半
	count = 0
	for i := range nums {
		if nums[i] == now {
			count++
		}
	}
	if count >= len(nums)/2+1 {
		return now
	}
	return -1

}
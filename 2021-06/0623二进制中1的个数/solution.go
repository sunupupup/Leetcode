func hammingWeight(num uint32) int {
	ret := 0
	for num > 0 {
		if num%2 == 1 {
			ret++
		}
		num /= 2
	}
	return ret
}
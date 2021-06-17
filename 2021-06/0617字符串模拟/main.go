func isNumber(s string) bool {
	//模拟法
	n := len(s)
	cs := []byte(s)
	idx := -1
	for i := 0; i < n; i++ {
		if cs[i] == 'e' || cs[i] == 'E' {
			if idx == -1 {
				idx = i
			} else {
				//重复出现e的情况，返回false
				return false
			}
		}
	}
	ans := true
	if idx != -1 {
		//如果有e，检查前后两个数
		ans = check(cs, 0, idx-1, false) && ans
		ans = check(cs, idx+1, n-1, true) && ans
	} else {
		//如果没有e
		ans = check(cs, 0, n-1, false) && ans
	}
	return ans
}

func check(cs []byte, start int, end int, mustInteger bool) bool {

	if start > end {
		return false
	}
	//检查并去除+ - 号
	if cs[start] == '+' || cs[start] == '-' {
		start++
	}
	hasDot := false
	hasNum := false
	//检查dot
	for i := start; i <= end; i++ {
		if cs[i] == '.' {
			if mustInteger || hasDot {
				return false
			}
			hasDot = true
		} else if cs[i] >= '0' && cs[i] <= '9' {
			hasNum = true
		} else {
			//如果还出现了其他符号，如 + - 号，就返回false
			return false
		}
	}
	return hasNum

}



// 栈 + 哈希，多层map按括号层次统计
func countOfAtoms(formula string) string {
	//i是记录当前遍历到的位置
	i, n := 0, len(formula)

	//解析出一个原子
	parseAtom := func() string {
		start := i
		i++
		for i < n && unicode.IsLower(rune(formula[i])) {
			i++ //扫描首字母后的小写字母
		}
		return formula[start:i]
	}

	parseNum := func() (num int) {
		if i == n || !unicode.IsDigit(rune(formula[i])) {
			//如果这里不是数字，因为有的原子只有一个，没有跟数字
			return 1
		}
		for ; i < n && unicode.IsDigit(rune(formula[i])); i++ {
			num = num*10 + int(formula[i]-'0')
		}
		return
	}

	//每有一对括号，就得有一层stk
	stk := []map[string]int{{}} //由哈希表组成的栈
	for i < n {
		if ch := formula[i]; ch == '(' {
			i++
			stk = append(stk, make(map[string]int))
		} else if ch == ')' {
			i++
			//解析出括号右侧的数字
			num := parseNum()
			atomNum := stk[len(stk)-1] //上层括号饿点原子map
			stk = stk[:len(stk)-1]     //这个括号里的统计完了，就丢弃掉这个map
			for atom, v := range atomNum {
				stk[len(stk)-1][atom] += v * num // 将括号内的原子数量乘上 num，加到上一层的原子数量中
			}
		} else {
			//解析出括号里的原子
			atom := parseAtom()
			num := parseNum()
			stk[len(stk)-1][atom] += num
		}
	}

	//对map进行排序
	type pair struct {
		atom string
		num  int
	}
	pairs := []pair{}
	for k, v := range stk[0] {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].atom < pairs[j].atom })

	ret := ""
	for i := range pairs {
		if pairs[i].num == 1 {
			ret += pairs[i].atom
		} else {
			ret += pairs[i].atom + strconv.Itoa(pairs[i].num)
		}
	}
	return ret
}
//简单的哈希表去重，然后加点简单的逻辑，就是有点体力活
func displayTable(orders [][]string) [][]string {
	//先统计处有多少种菜,多少个桌号  去重
	foods := make(map[string]int)
	tables := make(map[string]int)

	for i := range orders {
		foods[orders[i][2]] = 0
		tables[orders[i][1]] = 0
	}

	foodslice := make([]string, 0)
	for k, _ := range foods {
		foodslice = append(foodslice, k)
	}
	sort.Slice(foodslice, func(i, j int) bool { return foodslice[i] < foodslice[j] })
	for i := range foodslice {
		foods[foodslice[i]] = i //表明每个食物在结果的位置
	}

	tableslice := make([]string, 0)
	for k, _ := range tables {
		tableslice = append(tableslice, k)
	}
	sort.Slice(tableslice, func(i, j int) bool {
		a, _ := strconv.Atoi(tableslice[i])
		b, _ := strconv.Atoi(tableslice[j])
		return a < b
	})
	for i := range tableslice {
		tables[tableslice[i]] = i //表明每个桌子在结果的位置
	}

	n := len(tableslice)
	m := len(foodslice)
	//现在有n张桌子、m个菜
	//fmt.Println(n, "   " ,m)
	temp := make([][]string, n)
	for i := 0; i < n; i++ {
		temp[i] = make([]string, m+1)
		temp[i][0] = tableslice[i]
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			temp[i][j] = "0"
		}
	}
	//fmt.Println(foodslice)
	//fmt.Println(tableslice)

	for i := range orders {
		tableNum := tables[orders[i][1]]
		foodNum := foods[orders[i][2]]
		preNum, _ := strconv.Atoi(temp[tableNum][foodNum+1])
		temp[tableNum][foodNum+1] = strconv.Itoa(preNum + 1)
	}

	ret := make([][]string, 0)
	firstLine := []string{"Table"}
	firstLine = append(firstLine, foodslice...)
	ret = append(ret, firstLine)
	ret = append(ret, temp...)
	//fmt.Println(ret)

	return ret
}
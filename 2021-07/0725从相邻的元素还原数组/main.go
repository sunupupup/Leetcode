/*
1743. 从相邻元素对还原数组
存在一个由 n 个不同元素组成的整数数组 nums ，但你已经记不清具体内容。好在你还记得 nums 中的每一对相邻元素。

给你一个二维整数数组 adjacentPairs ，大小为 n - 1 ，其中每个 adjacentPairs[i] = [ui, vi] 表示元素 ui 和 vi 在 nums 中相邻。

题目数据保证所有由元素 nums[i] 和 nums[i+1] 组成的相邻元素对都存在于 adjacentPairs 中，存在形式可能是 [nums[i], nums[i+1]] ，
也可能是 [nums[i+1], nums[i]] 。这些相邻元素对可以 按任意顺序 出现。

返回 原始数组 nums 。如果存在多种解答，返回 其中任意一个 即可。
*/

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	arrs := make([]int, n)
	m := map[int][]int{}
	for _, v := range adjacentPairs {
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}
	//fmt.Println(m)
	//找到头尾
	t := []int{}
	for k, v := range m {
		if len(v) == 1 { //其实只要找到头就行了
			t = append(t, k)
		}
	}

	//从temp[0]开始放数字
	arrs[0] = t[0]
	pre := t[0]
	next := m[t[0]][0]
	for i := 1; i < n; i++ {
		arrs[i] = next
		if next == t[1] {
			break
		}
		temp := m[next]
		pre, next = next, helper(temp, pre)
	}

	return arrs
}

func helper(a []int, b int) int {
	if a[0] == b {
		return a[1]
	}
	return a[0]
}

//只要找到一个头就行了
func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	arrs := make([]int, n)
	m := map[int][]int{}
	for _, v := range adjacentPairs {
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}

	head := 0
	for k, v := range m {
		if len(v) == 1 {
			head = k
			break
		}
	}

	//从temp[0]开始放数字
	arrs[0] = head
	pre := head
	next := m[head][0]
	for i := 1; i < n; i++ {
		arrs[i] = next
		temp := m[next]
		pre, next = next, helper(temp, pre)
	}

	return arrs
}

func helper(a []int, b int) int {
	if len(a) != 2 {
		return -1
	}
	if a[0] == b {
		return a[1]
	}
	return a[0]
}
//肯定不是构造树
//如果是正常顺序的话  label / 2 就是父节点 ，只是现在父节点的顺序对称了一下
func pathInZigZagTree(label int) []int {
	//节点除二，就可以得到对称的父节点
	//先算出在第几层
	floor := 0
	for i := 1; ; i++ {
		if 1<<i-1 >= label {
			floor = i
			break
		}
	}

	//那么ret中会有 floor-1 个值
	ret := []int{label}
	now := label
	for i := floor - 1; i >= 1; i-- {
		now = getFather(i, now)
		ret = append(ret, now)
	}
	sort.Ints(ret)
	return ret
}

func getFather(floor int, now int) int {
	head := 1 << (floor - 1)
	tail := 1<<floor - 1
	temp := now / 2
	return head + (tail - temp)
}
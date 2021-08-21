package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算最小航行费用
 * @param input int整型二维数组 二维网格
 * @return int整型
 */
type A struct {
	x    int
	y    int
	cost int
}

func minSailCost(input [][]int) int {
	// write code here
	rows := len(input)
	cols := len(input[0])
	//宽搜

	//防止回头
	mask := make([][]int, rows)
	for i := 0; i < rows; i++ {
		mask[i] = make([]int, cols)
	}

	var getAround func(A) []A
	getAround = func(a A) []A {

		x := a.x
		y := a.y
		cost := a.cost

		ret := []A{}
		//向左
		if x-1 >= 0 {
			if mask[x-1][y] == 0 && input[x-1][y] != 2 {
				ret = append(ret, A{x - 1, y, cost + getcost(input[x-1][y])})
				mask[x-1][y] = 1
			}
		}
		//向右
		if x+1 <= cols-1 {
			if mask[x+1][y] == 0 && input[x+1][y] != 2 {
				ret = append(ret, A{x + 1, y, cost + getcost(input[x+1][y])})
				mask[x+1][y] = 1
			}
		}
		//向上
		if y-1 >= 0 {
			if mask[x][y-1] == 0 && input[x][y-1] != 2 {
				ret = append(ret, A{x, y - 1, cost + getcost(input[x][y-1])})
				mask[x][y-1] = 1
			}
		}
		//向下
		if y+1 <= rows-1 {
			if mask[x][y+1] == 0 && input[x][y+1] != 2 {
				ret = append(ret, A{x, y + 1, cost + getcost(input[x][y+1])})
				mask[x][y+1] = 1
			}
		}

		return ret
	}

	ret := -1
	arr := []A{{0, 0, 0}}
	for len(arr) != 0 {
		temp := []A{}
		for i := 0; i < len(arr); i++ {
			mask[arr[i].x][arr[i].y] = 1
			temp = append(temp, getAround(arr[i])...)
		}
		for i := 0; i < len(temp); i++ {
			if temp[i].x == cols-1 && temp[i].y == rows-1 {
				if ret == -1 {
					ret = temp[i].cost
				} else {
					ret = min(ret, temp[i].cost)
				}
			}
		}
		arr = temp
	}

	return ret
}

func getcost(a int) int {
	if a == 1 {
		return 1
	}
	if a == 0 {
		return 2
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

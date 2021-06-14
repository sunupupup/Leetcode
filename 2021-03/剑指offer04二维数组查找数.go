package main

//找到符合的行  二分查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	indexs := make([]int, 0)
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}

	for i, v := range matrix {
		if v[0] <= target && target <= v[m-1] {
			indexs = append(indexs, i)
		}
	}

	for _, index := range indexs {
		if binaryserch(matrix[index], target) {
			return true
		}
	}
	return false

}

func binaryserch(v []int, target int) bool {
	n := len(v)
	a, b := 0, n-1
	for a <= b {
		mid := (a + b) / 2
		if v[mid] == target {
			return true
		} else if v[mid] > target {
			b = mid - 1
		} else {
			a = mid + 1
		}
	}
	return false
}

//线性查找:
func findNumberIn2DArray_2(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}

	row, col := 0, m-1

	for row < n && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false

}

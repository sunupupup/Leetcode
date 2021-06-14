package main

//空间消耗为 O(1)的方法
//大致思路:
//1.先判断有没有0存在，没有的话直接返回，有的话将这0所在的行列全都置为0
//2.用上述记为0的 行列 记录其他行列有没有0

//改进方法:
//1.先记录第一行、第一列 有没有0
//2.遍历其他行列，若有0，则修改第一行第一列对应的那个值为 0
//3.遍历第一行第一列，根据0，将其他行列置为0
//4.若步骤1中判断第一行第一列有0，则将第一行第一列置为0，否则直接return
func setZeroes(matrix [][]int) {
	m := len(matrix)    //m行
	n := len(matrix[0]) //n列

	p, exit := hasZero(matrix)
	if exit == false {
		//没有零
		return
	} else {
		row := p[0] //所在行
		col := p[1] //所在列
		//有零，先把这这0所在的行列置零
		//列置0
		for i := 0; i < m; i++ {
			matrix[i][col] = 0
		}
		//行置0
		for i := 0; i < n; i++ {
			matrix[row][i] = 0
		}

		//记录其他0在哪，记录到这一行或者列中
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if matrix[i][j] == 0 {
					matrix[row][j] = 1
					matrix[i][col] = 1
				}
			}
		}

		//记录完了之后，遍历这个 行 和 列 ，将其他记录的行列置零
		//扫描 matrix[row] 行，有哪出有零，就把这列置零
		for i := 0; i < n; i++ {
			if matrix[row][i] == 1 {
				for j := 0; j < m; j++ {
					matrix[j][i] = 0
				}
			}
		}
		//扫描 col 列，哪有零，把这一行变为0
		for i := 0; i < m; i++ {
			if matrix[i][col] == 1 {
				for j := 0; j < n; j++ {
					matrix[i][j] = 0
				}
			}
		}
	}

}

func hasZero(matrix [][]int) ([2]int, bool) {
	m := len(matrix)    //m行
	n := len(matrix[0]) //n列
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				p := [2]int{i, j}
				return p, true
			}
		}
	}
	temp := [2]int{0, 0}
	return temp, false
}

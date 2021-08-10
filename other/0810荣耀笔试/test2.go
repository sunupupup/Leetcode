package main

import "fmt"

func main() {

	T := 0
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		m, i, n := 0, 0, 0
		fmt.Scan(&m, &i, &n)

		matrix := make([][]int, m)
		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		count := 0
		fmt.Scan(&count)
		for o := 0; o < count; o++ {
			a1, a2, a3 := 0, 0, 0
			fmt.Scan(&a1, &a2, &a3)
			matrix[a1][a2] = a3
			matrix[a2][a1] = a3
		}

		//找用户i的n度好友
		mask := make([]int, m)
		friends := []int{i}
		for z := 0; z < n; z++ {
			temp := []int{}
			for _, v := range friends {
				temp = append(temp, getFriend(&matrix, v, &mask)...)
			}
			friends = temp
		}
		fmt.Println(friends)
	}
}

func getFriend(m *[][]int, man int, mask *[]int) []int {
	ret := []int{}
	for i := 0; i < len(*m); i++ {
		if i != man && (*m)[man][i] > 0 && (*mask)[i] != 1 {
			ret = append(ret, i)
			(*mask)[i] = 1
		}
	}
	return ret
}

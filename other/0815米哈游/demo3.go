package main

import (
	"fmt"
)

type MyPoint struct {
	x int
	y int
}

func main() {
	T := 0
	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		lines := 0
		cols := 1
		fmt.Scan(&lines, &cols)
		t1, t2, t3, t4 := 0, 0, 0, 0
		fmt.Scan(&t1, &t2, &t3, &t4)
		p_start := MyPoint{t1 - 1, t2 - 1}
		p_end := MyPoint{t3 - 1, t4 - 1}

		ditu := make([][]byte, lines)
		//用mask防止回头
		mask := make([][]int, lines)

		for i := 0; i < lines; i++ {
			s := ""
			fmt.Scan(&s)
			ditu[i] = []byte(s)
			mask[i] = make([]int, cols)
		}

		//广度优先，找到公主

		//获取周围的点
		var get4MyPoints func(MyPoint) []MyPoint
		get4MyPoints = func(p MyPoint) []MyPoint {
			ret := []MyPoint{}
			if p.x+1 < cols {
				if mask[p.x+1][p.y] == 0 && ditu[p.x+1][p.y] != '#' {
					ret = append(ret, MyPoint{p.x + 1, p.y})
					mask[p.x+1][p.y] = 1
				}
			}
			if p.x-1 >= 0 {
				if mask[p.x-1][p.y] == 0 && ditu[p.x-1][p.y] != '#' {
					ret = append(ret, MyPoint{p.x - 1, p.y})
					mask[p.x-1][p.y] = 1
				}
			}
			if p.y+1 < lines {
				if mask[p.x][p.y+1] == 0 && ditu[p.x][p.y+1] != '#' {
					ret = append(ret, MyPoint{p.x, p.y + 1})
					mask[p.x][p.y+1] = 1
				}
			}
			if p.y-1 >= 0 {
				if mask[p.x][p.y-1] == 0 && ditu[p.x][p.y-1] != '#' {
					ret = append(ret, MyPoint{p.x, p.y - 1})
					mask[p.x][p.y-1] = 1
				}
			}
			return ret
		}

		arr := []MyPoint{p_start}
		steps := 0
		flag := false
		for len(arr) != 0 && !flag {

			//先检查从这个arr里面出发能不能直接将雷排了
			if check2(&arr, &p_end) {
				flag = true
				break
			}
			//开始广搜
			temp := []MyPoint{}
			for i := range arr {
				temp = append(temp, get4MyPoints(arr[i])...)
			}
			if check(&temp, &p_end) {
				flag = true
				break
			}
			arr = temp
			steps++
		}
		if flag {
			fmt.Println((t1 * t3) ^ steps ^ (t2 * t4))
		} else {
			fmt.Println(-1)
		}
	}
}

//检查这些点中有没有目标点
func check(ps *[]MyPoint, target *MyPoint) bool {
	for _, v := range *ps {
		if v.x == target.x && v.y == target.y {
			return true
		}
	}
	return false
}

func check2(ps *[]MyPoint, target *MyPoint) bool {
	for _, v := range *ps {
		if abs(v.x-target.x) <= 1 && abs(v.y-target.y) <= 1 {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

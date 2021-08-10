package main

import (
	"fmt"
	"sort"
)

func main1() {

	T := 0
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		n := 0
		fmt.Scan(&n)
		times := [][]int{}
		for i := 0; i < n; i++ {
			a, b := 0, 0
			fmt.Scan(&a, &b)
			times = append(times, []int{a, b})
		}

		fmt.Println(T)
		fmt.Println(n)

		//按照开头进行排序，最早的且时间区间最长的在前面
		sort.Slice(times, func(i, j int) bool {
			return times[i][0] < times[j][0] || (times[i][0] == times[j][0] && times[i][1] > times[j][1])
		})

		for _, v := range times {
			fmt.Println(v)
		}

		var dfs func(int, int) int
		dfs = func(index int, now int) int {
			ret := 0
			for i := index + 1; i < len(times); i++ {
				if times[i][0] >= now {
					ret = max(ret, dfs(i, times[i][1]))
				}
			}
			return times[index][1] - times[index][0] + ret
		}
		ret := 0
		l := len(times)
		for i := 0; i < l; i++ {
			ret = max(dfs(i, times[i][1]), ret)
		}
		fmt.Println(ret)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func main() {

	T := 0
	fmt.Scan(&T)

	for t := 0; t < T; t++ {

		n := 0
		fmt.Scan(&n)
		times := [][]int{}
		for i := 0; i < n; i++ {
			a, b := 0, 0
			fmt.Scan(&a, &b)
			times = append(times, []int{a, b})
		}

		fmt.Println(T)
		fmt.Println(n)

		//按照开头进行排序，最早的且时间区间最长的在前面
		sort.Slice(times, func(i, j int) bool {
			return times[i][0] < times[j][0] || (times[i][0] == times[j][0] && times[i][1] > times[j][1])
		})

		for _, v := range times {
			fmt.Println(v)
		}

		endtime := 8
		ret := 0
		for len(times) > 0 && endtime != 24 {
			//加上第一个区间的长度
			endtime = times[0][1]
			ret += times[0][1] - times[0][0]
			fmt.Println(times[0][1],times[0][0])
			//推迟后面所有的会议
			temp := [][]int{}
			for i := range times {
				if i != 0 {
					if times[i][0] < endtime {
						v := endtime - times[i][0]
						times[i][0] += v
						times[i][1] += v
					}
					if times[i][0] <= 24 && times[i][1] <= 24 {
						temp = append(temp, times[i])
					}
				}
			}
			sort.Slice(temp, func(i, j int) bool {
				return temp[i][0] < temp[j][0] || (temp[i][0] == temp[j][0] && temp[i][1] > temp[j][1])
			})
			times = temp
		}
		fmt.Println(ret)
	}

}
*/

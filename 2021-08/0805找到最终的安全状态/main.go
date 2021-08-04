/*
802. 找到最终的安全状态
在有向图中，以某个节点为起始节点，从该点出发，每一步沿着图中的一条有向边行走。
如果到达的节点是终点（即它没有连出的有向边），则停止。

对于一个起始节点，如果从该节点出发，无论每一步选择沿哪条有向边行走，最后必然在有限步内到达终点，
则将该起始节点称作是 安全 的。

返回一个由图中所有安全的起始节点组成的数组作为答案。答案数组中的元素应当按 升序 排列。

该有向图有 n 个节点，按 0 到 n - 1 编号，其中 n 是 graph 的节点数。图以下述形式给出：
graph[i] 是编号 j 节点的一个列表，满足 (i, j) 是图的一条有向边。

*/
package main

import "fmt"

//官方：
//深搜找环 + 三色标记
func eventualSafeNodes(graph [][]int) []int {

	color := make([]int, len(graph))

	//从A点出发，发现有环了，说明这个环上所有点有环
	//发现无环，说明深搜所有点无环

	var safe func(now int) bool
	safe = func(now int) bool {
		if color[now] > 0 {
			return color[now] == 2
		}

		color[now] = 1
		for _, v := range graph[now] {
			if !safe(v) {
				return false
			}
		}
		color[now] = 2
		return true
	}

	ret := []int{}
	for i := 0; i < len(graph); i++ {
		if safe(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

//我的做法：深搜找环
//找有向图中的环
//超时：105 / 112 个通过测试用例
func eventualSafeNodes3(graph [][]int) []int {

	//graph就相当于是一个map了

	n := len(graph)

	//找环，环上所有的节点都不在结果之内
	mask := make([]int, n)

	var dfs func(int, int, *[]int) bool
	dfs = func(root int, now int, temp *[]int) bool {
		for _, v := range graph[now] {

			//如果之前已经访问到这个节点，表明有环
			if (*temp)[v] == 1 {
				return true
			}

			//到了一个环上的节点  或者  到了一个确定没有环的节点
			if mask[v] == 1 {
				return true
			} else if mask[v] == 0 && v < root {
				return false
			}

			(*temp)[v] = 1
			if dfs(root, v, temp) {
				return true
			}
			(*temp)[v] = 0

		}
		return false
	}

	for i := 0; i < n; i++ {
		//深搜
		if mask[i] == 0 {
			//fmt.Println(i)
			temp := make([]int, n)
			temp[i] = 1
			//如果成环了，将路线上所有的点记录到mask中，记为1
			if dfs(i, i, &temp) {
				for i, v := range temp {
					if v == 1 {
						mask[i] = 1
					}

				}
			}
		}
	}

	ret := []int{}
	for i, v := range mask {
		if v == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	a := [][]int{
		{1, 2, 3, 4},
		{1, 2},
		{3, 4},
		{0, 4},
		{},
	}

	ret := eventualSafeNodes2(a)
	fmt.Println(ret)
}

func eventualSafeNodes2(graph [][]int) []int {

	//graph就相当于是一个map了

	n := len(graph)

	//找环，环上所有的节点都不在结果之内
	mask := make([]int, n)

	var dfs func(int, int, *map[int]struct{}) bool
	dfs = func(root int, now int, temp *map[int]struct{}) bool {
		for _, v := range graph[now] {
			if _, ok := (*temp)[v]; ok {
				return true
			}
			if mask[v] != 1 {
				(*temp)[v] = struct{}{}
				if dfs(root, v, temp) {
					return true
				}
				delete(*temp, v)
			} else {
				return false
			}
		}
		return false
	}

	var dfs2 func(int, *map[int]struct{}, *[]int)
	dfs2 = func(now int, temp *map[int]struct{}, temp2 *[]int) {
		for _, v := range graph[now] {

			if (*temp2)[v] == 1 {
				continue
			}
			if mask[v] == -1 {
				return
			}
			(*temp)[v] = struct{}{}
			(*temp2)[v] = 1
			dfs2(v, temp, temp2)

		}
	}

	for i := 0; i < n; i++ {
		//深搜
		if mask[i] == 0 {
			if len(graph[i]) == 0 {

				mask[i] = -1
				continue
			}
			temp := map[int]struct{}{i: struct{}{}}
			//如果成环了，将路线上所有的点记录到mask中，记为1,明确不成环了，记为-1
			if dfs(i, i, &temp) {
				for k, _ := range temp {
					mask[k] = 1
				}
			} else {
				temp2 := make([]int, n)
				temp2[i] = 1
				dfs2(i, &temp, &temp2)
				for k, _ := range temp {
					mask[k] = -1
				}
			}
		}
	}

	ret := []int{}
	//fmt.Println(mask)

	for i, v := range mask {
		if v == -1 {
			ret = append(ret, i)
		}
	}
	return ret
}

package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g) //升序排列
	//sort.Sort(sort.Reverse(sort.IntSlice(g)))  降序
	sort.Ints(s)
	gn := len(g)
	sn := len(s)
	i := 0
	j := 0
	for i < gn && j < sn {
		if s[j] >= g[i] {
			i++
			j++
		} else {
			j++
		}

	}
	return i
}

func test() {

	arrayInt := []int{3, 1, 2, 5, 4}

	arrayFloat := []float64{3.2, 1.8, 1.9, 2.2, 4.3}

	arrayString := []string{"abc", "ab", "bc"}

	// 升序
	sort.Ints(arrayInt)
	sort.Float64s(arrayFloat)
	sort.Strings(arrayString)

	fmt.Println("sort int:", arrayInt)
	fmt.Println("sort float:", arrayFloat)
	fmt.Println("sort ", arrayString)

	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(arrayInt)))
	sort.Sort(sort.Reverse(sort.Float64Slice(arrayFloat)))
	sort.Sort(sort.Reverse(sort.StringSlice(arrayString)))

	fmt.Println("After reversed: ")

	fmt.Println("sort int:", arrayInt)
	fmt.Println("sort float:", arrayFloat)
	fmt.Println("sort ", arrayString)

}

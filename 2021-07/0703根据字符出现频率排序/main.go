package main

import (
	"bytes"
	"sort"
)

type T struct {
	a byte
	b int
}

func frequencySort(s string) string {
	sb := []byte(s)
	temp := [256]int{0}
	for _, b := range sb {
		temp[int(b)]++
	}

	//统计出现的字符和出现的次数
	counts := make([]T, 0)
	for i, v := range temp {
		if v != 0 {
			counts = append(counts, T{byte(i), v})
		}
	}

	//对上面统计的，按照出现的次数进行排序
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].b > counts[j].b
	})

	//预留 len(s) 的空间，减少内存的分配
	ret := make([]byte, 0, len(s))
	for _, v := range counts {
		for i := 0; i < v.b; i++ {
			ret = append(ret, v.a)
		}
	}

	return string(ret)
}

//因为直到每个字符出现的上限频次
//所以可以使用桶排序
//定义 maxFrquency 个桶
func frequencySort2(s string) string {
	//因为单个字符出现的频率是有上限的
	n := len(s)
	m := make(map[byte]int)
	maxFrequency := 0
	for i := range s {
		m[s[i]]++
		maxFrequency = max(maxFrequency, m[s[i]])
	}

	temp := make([][]byte, maxFrequency+1)
	for k, v := range m {
		temp[v] = append(temp[v], k)
	}

	ret := make([]byte, 0, n)
	for i := maxFrequency; i >= 0; i-- {
		for _, v := range temp[i] {
			ret = append(ret, bytes.Repeat([]byte{v}, i)...)
		}
	}

	return string(ret)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
451. 根据字符出现频率排序
给定一个字符串，请将字符串里的字符按照出现的频率降序排列。

示例 1:
输入:"tree"
输出:"eert"

解释:
'e'出现两次，'r'和't'都只出现一次。
因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
*/

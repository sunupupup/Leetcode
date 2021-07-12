/*
275. H 指数 II
给定一位研究者论文被引用次数的数组（被引用次数是非负整数），数组已经按照 升序排列 。编写一个方法，计算出研究者的 h 指数。

h 指数的定义: “h 代表“高引用次数”（high citations），
一名科研人员的 h 指数是指他（她）的 （N 篇论文中）总共有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）"

*/
package main

import "sort"

func hIndex(citations []int) int {
	left, right := 0, citations[len(citations)-1]+1

	helper := func(h int) bool {
		count := 0
		for i := range citations {
			if citations[i] >= h {
				count++
			}
		}
		return count >= h
	}

	for left < right {
		mid := (right-left)/2 + left
		if helper(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

//官方：由于数组 citations 已经按照升序排序，因此可以使用二分查找。
//设查找范围的初始左边界 left 为 0,
//初始右边界 right 为 n-1，
//其中 n 为数组 citations 的长度。
//每次在查找范围内取中点mid，则有 n−mid 篇论文被引用了至少 citations[mid] 次。
//如果在查找过程中满足citations[mid]≥n−mid，则移动右边界right，否则移动左边界 left。

func hIndex2(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(x int) bool { return citations[x] >= n-x })
}

package main

import (
	"container/heap"
	"fmt"
)

type myHeap [][2]int

func (mh myHeap) Len() int { return len(mh) }
func (mh myHeap) Less(i, j int) bool {
	return mh[i][0] < mh[j][0] || (mh[i][0] == mh[j][0] && mh[i][1] < mh[j][1])
}
func (mh *myHeap) Swap(i, j int) { (*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i] }
func (mh *myHeap) Pop() interface{} {
	ret := (*mh)[len(*mh)-1]
	(*mh) = (*mh)[:len(*mh)-1]
	return ret
}
func (mh *myHeap) Push(x interface{}) {
	item := x.([2]int)
	(*mh) = append((*mh), item)
}

func main() {
	h := myHeap{{1, 2}, {2, 2}, {4, 1}, {0, 3}, {4, 6}, {9, 3}}

	heap.Init(&h)

	heap.Push(&h, [2]int{10, 10})

	for len(h) > 0 {
		fmt.Println(heap.Pop(&h).([2]int))
	}

}

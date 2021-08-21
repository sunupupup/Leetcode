package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	temp := ""
	s := ""
	for {
		_, err := fmt.Scan(&temp)
		if err != nil {
			break
		}
		s += temp
	}
	arr1 := strings.Split(s, ",")
	arr := []int{}
	for i := 0; i < len(arr1); i++ {
		num, _ := strconv.Atoi(arr1[i])
		arr = append(arr, num)
	}

	start := 0
	end := len(arr) - 1

	maxRet := 0
	var helper func(start, end, score1, score2 int)
	helper = func(start, end, score1, score2 int) {
		if start > end {
			maxRet = max(maxRet, score1)
			return
		}
		if start == end {
			maxRet = max(maxRet, score1+arr[end])
			return
		}
		if start+1 < len(arr) {
			helper(start+2, end, score1+arr[start], score2+arr[start+1])
		}
		if end-1 >= 0 {
			helper(start, end-2, score1+arr[end], score2+arr[end-1])
		}
		helper(start+1, end-1, score1+arr[start], score2+arr[end])
		helper(start+1, end-1, score1+arr[end], score2+arr[start])

	}

	helper(start, end, 0, 0)

	fmt.Println(maxRet)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

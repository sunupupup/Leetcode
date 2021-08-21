package main

import (
	"fmt"
)

func main() {

	fmt.Println(string(findKthBit(4, 11)))
}

func findKthBit(n int, k int) byte {

	temp := make([]string, n+1)
	temp[0] = ""
	temp[1] = "a"

	for i := 2; i <= n; i++ {
		bs := []byte{}
		bs = append(bs, []byte(temp[i-1])...)
		bs = append(bs, 'a'+byte(i-1))
		bs = append(bs, []byte(reverseAndInvert(temp[i-1]))...)
		temp[i] = string(bs)
	}
	return temp[n][k-1]
}
func reverseAndInvert(s string) string {
	bs := []byte(s)
	n := len(s)
	bs_ret := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		bs_ret[n-1-i] = 'z' - bs[i] + 'a'
	}
	return string(bs_ret)
}

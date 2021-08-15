/*
小美现在有一个字符串，小美现在想知道能不能通过在字符串的尾端增加若干字符使得整个字符串变成一个回文串。

回文串的定义：若一个字符串，对他正序遍历和倒序遍历得到的结果是完全一致的，就称它是一个回文串。例如 abcba 就是一个回文串，因为无论正序还是倒序都是一样的。


对于字符串 abaaca，显然在该字符串末尾继续补上三个字符 aba 就可以构成 abaacaaba，就可以把原字符串变成回文串。换句话说，最少补上三个字符。

你的任务就是找到使得原来的字符串变成回文串所需要的最少字符数量。

* 本题数据保证没有空串，因此不考虑空串是否为回文串。

保证输入的字符串仅包含小写字母。

输入描述
一行一个字符串，代表小美交给你的字符串。

输出描述
一行一个整数，表示将小美给出的字符串变成回文字符串所需要添补的最少字符数量。
*/

package main

import "fmt"

func test2() {

	s := ""
	fmt.Scan(&s)

	var helper func(int, int) int
	helper = func(i1, i2 int) int {

		ret := 0
		for s[i1] == s[i2] && i2 > i1 {
			ret += 2
			i1++
			i2--
		}
		if i1==i2 {
			return ret+1
		}else if i2==i1-1{
			return ret
		}
		return -1
	}

	//检查这个s，以最后一个字符结尾的最长回文子串
	n := len(s)
	record := []int{}
	tail := s[n-1]
	for i := 0; i < n-1; i++ {
		if s[i] == tail {
			record = append(record, i)
		}
	}
	ret := 1
	for i := 0; i < len(record); i++ {
		//计算以 record[i]开头，n-1结尾的字符串是不是回文串
		ret = max(ret, helper(record[i], n-1))
	}

	fmt.Println(n - ret)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func main() {


	s := ""
	fmt.Scan(&s)

	//检查这个s，以最后一个字符结尾的最长回文子串
	n := len(s)
	ret := 1
	for i := n - 1; i >= n/2; i-- {
		ret = max(ret, helper1(s, i, i))
		ret = max(ret, helper1(s, i-1, i))
	}
	fmt.Println(n - ret)
}

//start和end是扩散的位置
func helper1(s string, start int, end int) int {
	n := len(s)
	ret := -1
	for s[start] == s[end] {
		ret += 2
		if end == n-1 {
			return ret
		}
		start--
		end++
	}
	return ret
}

func helper2(s string, start int, end int) int {
	n := len(s)
	ret := 0
	for s[start] == s[end] {
		ret += 2
		if end == n-1 {
			return ret
		}
		start--
		end++
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

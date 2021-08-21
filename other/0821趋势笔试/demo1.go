package main

import (
	"fmt"
	"strings"
)

func main() {
	templateString := ""
	fmt.Scan(&templateString)

	temp := ""
	s := ""
	for {
		_, err := fmt.Scan(&temp)
		if err != nil {
			break
		}
		s += temp
	}
	arr := strings.Split(s[1:len(s)-1], ",")

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i][1 : len(arr[i])-1]
	}
	//先判断 * 的位置
	//如果是第四种情况
	if templateString == "*" {
        pln(arr)
        return 
	}

	//找 * 的位置
	index := -1
	ss := strings.Split(templateString, ".")
	for i := 0; i < len(ss); i++ {
		if ss[i] == "*" {
			index = i
		}
	}

	ret := []string{}
	if index == -1 {
		for i := 0; i < len(arr); i++ {
			if arr[i] == templateString {
				ret = append(ret, arr[i])
			}
		}
		pln(ret)
        return 
	}
	//如果出现在第一个，只需要匹配后面的所有字符
	if index == 0 {
		for i := 0; i < len(arr); i++ {
			if helper1(templateString, arr[i]) {
				ret = append(ret, arr[i])
			}
		}
		pln(ret)
        return 
	}

	for i := 0; i < len(arr); i++ {
		if helper2(templateString, arr[i],index) {
			ret = append(ret, arr[i])
		}
	}
	pln(ret)

}

func pln(arr []string) {
    if len(arr)==0 {
        fmt.Print("Not Match")
        return 
    }
	for i := 0; i < len(arr); i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
	}
}

func helper1(s1, s2 string) bool {
	//从后往前匹配
	arr_s1 := strings.Split(s1, ".")
	arr_s2 := strings.Split(s2, ".")


	if len(arr_s2) <= len(arr_s1) {
		return false
	}

	for i := len(arr_s1) - 1; i >= 1; i-- {
        if arr_s1[i] != arr_s2[i + len(arr_s2)-len(arr_s1)] {
			return false
		}
	}
	return true
}

//* 在中间的时候
func helper2(s1, s2 string, index int) bool {
	//从后往前匹配
	arr_s1 := strings.Split(s1, ".")
	arr_s2 := strings.Split(s2, ".")
	if len(arr_s2) != len(arr_s1) {
		return false
	}
	for i := len(arr_s1) - 1; i >= 0; i-- {
		if arr_s1[i] != arr_s2[i] && i != index {
			return false
		}
	}
	return true
}

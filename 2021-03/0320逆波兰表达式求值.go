package main

import "strconv"

func evalRPN(tokens []string) int {
	n := len(tokens)
	index := -1
	stk := make([]int, n)
	for i := 0; i < n; i++ {
		token := tokens[i]
		if isDigit(token) {
			//如果是数字
			index++
			stk[index], _ = strconv.Atoi(token)
		} else {
			//如果是运算符
			switch token {
			case "+":
				index--
				stk[index] += stk[index+1]
			case "-":
				index--
				stk[index] -= stk[index+1]
			case "*":
				index--
				stk[index] *= stk[index+1]
			case "/":
				index--
				stk[index] /= stk[index+1]
			}
		}
	}
	return stk[index]
}

func isDigit(s string) bool {
	return s != "+" && s != "-" && s != "*" && s != "/"
}

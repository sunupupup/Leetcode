//输入包括两个正整数a,b(1 <= a, b <= 10^9),输入数据有多组, 如果输入为0 0则结束输入
package main

import "fmt"

func tes3() {
	var a, b int
	for {
		fmt.Scan(&a, &b)
		if a == b && a == 0 {
			break
		} else {
			fmt.Println(a + b)
		}
	}
}

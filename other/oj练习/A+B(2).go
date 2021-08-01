package main

import "fmt"

/*
输入第一行包括一个数据组数t(1 <= t <= 100)
接下来每行包括两个正整数a,b(1 <= a, b <= 10^9)

2
1 5
10 20
*/

func test2_1() {
    var t, a, b int
    fmt.Scan(&t) // 
    for t > 0 {
        fmt.Scan(&a, &b)
        fmt.Println(a + b)
        t--
    }
}

//业务场景还要考虑 err的情况
func test2_2() {
    var t, a, b int
    if _, err := fmt.Scan(&t); err != nil {
        // 处理错误
        return
    } 
    for t > 0 {
        if _, err := fmt.Scan(&a, &b); err != nil {
        	// 处理错误
        	return
    	} 
        fmt.Println(a + b)
        t--
    }
}

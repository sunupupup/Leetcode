//输入包括两个正整数a,b(1 <= a, b <= 10^9),输入数据包括多组。
package main  // 包名， 必需的，因为这里就是相当于一个go文件

import (  // 调用的包
    "fmt"
    "io"
)

func test1() {
    var a, b int
    for { // 循环获取输入
        if _, err := fmt.Scan(&a,&b); err != io.EOF {
            fmt.Println(a + b)
        } else {
            break
        }
    }
}
/*
输入数据包括多组。
每组数据一行,每行的第一个整数为整数的个数n(1 <= n <= 100), n为0的时候结束输入。
接下来n个正整数,即需要求和的每个正整数。
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, num int
	for {
		fmt.Scan(&n) //正常先扫描个数
		if n == 0 {
			break
		}
		sum := 0
		for n > 0 {
			fmt.Scan(&num)
			sum += num
			n--
		}
		fmt.Println(sum)
	}
}

// 当然这个题可以使用bufio包来做，也就是读入一行数据再进行处理

func test4_2() {
	input := bufio.NewScanner(os.Stdin) //创建并返回一个从os.Stdin读取数据的Scanner
	for input.Scan() {
		// Scan方法获取当前位置的token（该token可以通过Bytes或Text方法获得），
		// 并让Scanner的扫描位置移动到下一个token。
		// 当扫描因为抵达输入流结尾或者遇到错误而停止时，本方法会返回false
		nums := strings.Split(input.Text(), " ") //分割字符串	获取当前行
		if nums[0] == "0" {                      // 判断是否结束
			break
		}
		res := 0
		for i := 1; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i]) // 字符串转数字
			res += num
		}
		fmt.Println(res)
	}
}

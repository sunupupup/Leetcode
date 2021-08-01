/*
多个测试用例，每个测试用例一行。

每行通过空格隔开，有n个字符，n＜100


对于每组测试用例，输出一行排序过的字符串，每个字符串通过空格隔开
示例1
输入
a c bb
f dddd
nowcoder
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := strings.Split(input.Text(), " ") // 字符串分割
		sort.StringSlice.Sort(str)              // 字符串排序，这里借助了sort.StringSlice方法来实现
		fmt.Println(strings.Join(str, " "))     // 通过Join来连接， 也可以手动连接
	}
}

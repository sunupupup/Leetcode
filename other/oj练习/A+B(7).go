/*
输入数据有多组, 每行表示一组输入数据。
每行不定有n个整数，空格隔开。(1 <= n <= 100)。
*/

// 这个直接用bufio做好，因为这题没有指定数组的大小了，
// 这种情况用fmt.Scan我还没想到，有做出来的大佬们可以告诉我一下！感谢
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := strings.Split(input.Text(), " ")
		res := 0
		for i := 0; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])
			res += num
		}
		fmt.Println(res)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//给定如下IP地址和网段，编码判断ip地址是否在网段内
func test1() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if len(input.Text()) < 1 {
			break
		}
		str := strings.Split(input.Text(), `"`)
		ip := str[1]
		wangduan := str[3]

		wangduan_slice := strings.Split(wangduan, "/")
		mask := wangduan_slice[0]
		numStr := wangduan_slice[1]

		ip4 := strings.Split(ip, ".")
		ipNums := []int{}
		for i := range ip4 {
			temp, _ := strconv.Atoi(ip4[i])
			ipNums = append(ipNums, temp)
		}

		mask4 := strings.Split(mask, ".")
		maskNums := []int{}
		for i := range ip4 {
			temp, _ := strconv.Atoi(mask4[i])
			maskNums = append(maskNums, temp)
		}

		num, _ := strconv.Atoi(numStr)

		//校验前num位
		ipAll := ipNums[0]<<24 + ipNums[1]<<16 + ipNums[2]<<8 + ipNums[3]
		maskAll := maskNums[0]<<24 + maskNums[1]<<16 + maskNums[2]<<8 + maskNums[3]

		fmt.Println(ipAll>>(32-num) == maskAll>>(32-num))
	}
}

//给定如下IP地址和网段，编码判断ip地址是否在网段内
/*
func main() {

	for {

		ip_temp := ""
		_, err := fmt.Scan(&ip_temp)
		if err == io.EOF {
			break
		}
		fmt.Scan(&ip_temp)
		//ip := ip_temp[1 : len(ip_temp)-1]

		wangduan := ""
		//fmt.Scan(&wangduan)
		fmt.Scan(&wangduan)

		fmt.Println(ip_temp)
		fmt.Println(wangduan)
	}

}

*/

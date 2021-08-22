package main

import "fmt"

func main() {

	m, n := 0, 0
	fmt.Scan(&m, &n)
	s := ""
	fmt.Scan(&s)

	bs := []byte(s)
	to_find := n //表示还需要找这么多
	has_find := 0
	temp := []byte{}
	end := -1
	//找now到now + m-n
	for to_find > n-(end+1) {
		index := end + 1
		b := bs[index]
		for i := end + 1; i <= end+1+(m-n+1+has_find); i++ {
			if bs[i] > b {
				b, index = bs[i], i
			}
		}
		end = index
		to_find--
		has_find++
		temp = append(temp, b)
	}
	temp = append(temp, bs[end+1:]...)
	fmt.Println(string(temp))
}

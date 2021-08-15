package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func test() {

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if len(input.Text()) < 1 {
			break
		}
		s := input.Text()[1 : len(input.Text())-1]
		temp, _ := time.Parse("2006-01-02 15:04:05", s)
		fmt.Println(temp.Unix() - 1617250332)
	}

}

func main() {

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if len(input.Text()) < 1 {
			break
		}
		s := input.Text()[1 : len(input.Text())-1]
		temp, _ := time.Parse("2006-01-02 15:04:05", s)
		fmt.Println(temp.Unix() - 28800)
	}

}

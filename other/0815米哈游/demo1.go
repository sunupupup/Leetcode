package main

import "fmt"

func demo1() {

	n := 0
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		s := ""
		fmt.Scan(&s)
		bs := []byte(s)
		count := 0
		for _, b := range bs {
			if b == 'a' {
				count++
			} else {
				count--
			}
			if count < 0 {
				fmt.Println("NO")
                break
			}
		}
		if count == 0 {
			fmt.Println("YES")
        }else if count>  0 {
            fmt.Println("NO")
        }

	}

}

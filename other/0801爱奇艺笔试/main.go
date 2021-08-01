package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func test1() {

	var s string
	fmt.Scan(&s)
	fmt.Println(s)
	numsStr := strings.Split(s, ",")
	count := len(numsStr)
	last := strings.Split(numsStr[count-1], ":")
	numsStr[count-1] = last[0]
	var size int
	size, _ = strconv.Atoi(last[1])

	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i], _ = strconv.Atoi(numsStr[i])
	}

	var f func(int) int
	f = func(start int) (ret int) {
		for i := 0; i < size; i++ {
			ret += nums[i+start]
		}
		return
	}

	preSum := 0
	var p float64
	p = -100000000
	for i := 0; i <= count-size; i++ {
		if i == 0 {
			preSum = f(0)
		} else {
			temp := f(i)
			p = max(p, float64(temp-preSum)/float64(preSum))
			preSum = temp
		}
	}
	p = p * 100
	fmt.Printf("%.2f", p)
	fmt.Println("%")

}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func test2() {
	var a string
	fmt.Scan(&a)
	bs := []byte(a)
	a = string(bs[1 : len(a)-1])
	rainsStr := strings.Split(a, ",")

	rains := make([]int, len(rainsStr))
	for i, v := range rainsStr {
		rains[i], _ = strconv.Atoi(v)
	}

	//记录可以抽水的次数

	count := 0
	for i := 0; i < len(rains); i++ {
		//先找有多少湖
		if rains[i] > count {
			count = rains[i]
		}
	}

	//所以湖的数量为count个
	//将消耗的机会，内容填进临时数组
	temp := []int{}
	mask := make([]int, count+1)
	k := 0
	for _, v := range rains {
		if v == 0 {
			k++
		} else {
			if mask[v] == 0 {
				mask[v] = 1
			} else {
				//利用前面抽水的机会
				if k > 0 {
					k--
					temp = append(temp, v) //抽v的水
				} else {
					fmt.Print("[]")
					break
				}
			}
		}
	}

	index := 0
	ret := make([]int, len(rains))
	for i := 0; i < len(rains); i++ {
		if rains[i] > 0 {
			ret[i] = -1
		} else {
			ret[i] = temp[index]
			if index < len(temp)-1 {
				index++
			}
		}
	}
	fmt.Print("[")
	s := ""
	for i := 0; i < len(ret); i++ {
		s += "," + strconv.Itoa(ret[i])
	}
	fmt.Print(s[1:])
	fmt.Print("]")
}

/*
var m sync.Map
var count int

func main() {
	//能否一次性读取呢？
	var s string
	fmt.Scan(&s)
	lines := strings.Split(s, ",")
	all := len(lines)

	m = sync.Map{}
	count = 0

	go func() {
		for i := 1; i <= all; i++ {
			line := lines[i-1]
			go func(i int) {
				m.Store(i, strings.ToUpper(line))
			}(i)
		}
	}()

	for i := 1; i <= all; i++ {
		if i == 1 {
			for {
				if v, ok := m.Load(i); !ok {
					continue
				} else {
					fmt.Print(v.(string))
					break
				}
			}
		} else {
			for {
				if v, ok := m.Load(i); !ok {
					continue
				} else {
					fmt.Print(",")
					fmt.Print(v.(string))
					break
				}
			}
		}
	}
}
*/

var mu sync.Mutex
var m map[int]string
var count int

func main() {
	//能否一次性读取呢？
	var s string
	fmt.Scan(&s)
	lines := strings.Split(s, ",")
	all := len(lines)

	m = map[int]string{}
	count = 0

	go func() {
		for i := 1; i <= all; i++ {
			line := lines[i-1]
			go func(i int) {
				m[i] = strings.ToUpper(line)
			}(i)
		}
	}()

	for i := 1; i <= all; i++ {
		if i == 1 {
			for {

				if v, ok := m[i]; !ok {
					continue
				} else {
					fmt.Print(v)
					break
				}
			}
		} else {
			for {
				if v, ok := m[i]; !ok {
					continue
				} else {
					fmt.Print(",")
					fmt.Print(v)
					break
				}
			}
		}
	}
}

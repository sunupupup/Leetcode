/*
题目描述：
二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。每个 LED 代表一个 0 或 1，最低位在右侧。
给你一个整数 turnedOn ，表示当前亮着的 LED 的数量，返回二进制手表可以表示的所有可能时间。你可以 按任意顺序 返回答案。
*/

//做法1：递归，从 固定数组中 取n个数
func readBinaryWatch(turnedOn int) (ret []string) {

	if turnedOn > 8 {
		return ret
	}
	hour := 0
	for hour <= turnedOn && hour <= 3 {
		hourString := getString(hour, "hour")
		mintueString := getString(turnedOn-hour, "minute")
		for i := 0; i < len(hourString); i++ {
			for j := 0; j < len(mintueString); j++ {
				if hourString[i] == "" || mintueString[j] == "" {
					continue
				}
				ret = append(ret, hourString[i]+":"+mintueString[j])
			}
		}

		hour++
	}

	return
}

func getString(n int, mode string) (ret []string) {
	if n == 0 {
		if mode == "hour" {
			return []string{"0"}
		} else if mode == "minute" {
			return []string{"00"}
		}
	}
	// 从 1 2 4 8中选出n个数的和
	var choose []int
	if mode == "hour" {
		choose = []int{8, 4, 2, 1}
	} else if mode == "minute" {
		choose = []int{32, 16, 8, 4, 2, 1}
	}

	//然后从choose中挑选出n个数，变成字符串返回 还要看考虑 minute的0-9多加上一个 0
	temp := make([]int, 0)
	dfs(choose, 0, n, 0, &temp)

	ret = make([]string, len(temp))
	for i := 0; i < len(ret); i++ {

		if mode == "minute" && (temp[i] >= 0 && temp[i] <= 9) {
			ret[i] = "0" + strconv.Itoa(temp[i])
		} else {
			if (mode == "minute" && temp[i] > 59) || (mode == "hour" && temp[i] >= 12) {
				continue
			}
			ret[i] = strconv.Itoa(temp[i])
		}
	}

	return ret
}

func dfs(choose []int, index int, count int, add int, temp *[]int) {

	if count == 0 || index == len(choose) {
		if count == 0 {
			*temp = append((*temp), add)
		}
		return
	}

	dfs(choose, index+1, count, add, temp)                 //不选  index处的数
	dfs(choose, index+1, count-1, add+choose[index], temp) //选    index处的数

}



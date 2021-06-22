/*
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

输入： "aac"
输出： ["aac","aca","caa"]
*/

func permutation(s string) []string {
	//abc  a+bc组合  b+ac组合  c+ab组合
	ret := helper(s)

	//利用map去重
	m := make(map[string]struct{}, 0)
	for _, v := range ret {
		m[v] = struct{}{}
	}

	result := make([]string, 0)
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}

func helper(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	ret := make([]string, 0)
	ss := []byte(s)
	for i := 0; i < len(s); i++ {
		//剩下的字符 可以组成的[]string
		leave := subStr(s, i)
		groups := helper(leave)
		for j := 0; j < len(groups); j++ {
			ret = append(ret, string(ss[i])+groups[j])
		}
	}
	return ret
}

//把s中下标为i的字母去掉，返回一个新的string
func subStr(s string, index int) string {
	ss := []byte(s)
	ret := make([]byte, 0, len(s)-1)
	for i := 0; i < len(s); i++ {
		if i == index {
			continue
		}
		ret = append(ret, ss[i])
	}
	return string(ret)
}






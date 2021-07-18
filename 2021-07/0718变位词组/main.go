/*
面试题 10.02. 变位词组
编写一种方法，对字符串数组进行排序，将所有变位词组合在一起。变位词是指字母相同，但排列不同的字符串。

注意：本题相对原题稍作修改

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
*/

//计算出每个str的 [26]int
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	a := [26]int{}
	for _, str := range strs {
		a = [26]int{}
		bs := []byte(str)
		for _, b := range bs {
			a[int(b-'a')]++
		}
		if _, ok := m[a]; !ok {
			m[a] = make([]string, 0)
		}
		m[a] = append(m[a], str)
	}
	ret := [][]string{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

//也可以对每个str的byte排序
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(a, b int) bool { return bs[a] < bs[b] })
		newstr := string(bs)
		if _, ok := m[newstr]; !ok {
			m[newstr] = make([]string, 0)
		}
		m[newstr] = append(m[newstr], str)
	}

	ret := [][]string{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}
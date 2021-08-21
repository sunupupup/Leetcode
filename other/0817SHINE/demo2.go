package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 从haystack找到第一个needle的位置
 *
 * @param haystack string字符串
 * @param needle string字符串
 * @return int整型
 */
func strStr(haystack string, needle string) int {
	// write code here
	if len(needle) == 0 || len(haystack) > len(needle) {
		return 0
	}
	ret := -1

	bs := []byte(haystack)
	b := []byte(needle)[0]
	indexs := []int{}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if bs[i] == b {
			indexs = append(indexs, i)
		}
	}
	n := len(needle)
	for _, index := range indexs {
		if string(bs[index:index+n]) == needle {
			return index
		}

	}

	return ret
}

//递归，没想过优化的地方

var length int = 0

func maxLength(arr []string) int {
	length = 0
	//递归 + 回溯
	mask := [26]int{0}
	for i := 0; i < len(arr); i++ {
		helper(i, arr, mask, 0)
	}
	return length
}

func helper(index int, arr []string, mask [26]int, count int) { //每次mask传的都是拷贝，所以无需回溯
	if index >= len(arr) {
		length = max(length, count)
		return
	}
	//将这个 index 的 string放到mask里
	temp := 0
	bytes := []byte(arr[index])
	for _, b := range bytes {
		mask[int(b-'a')]++
		temp++
		//如果有重复，就返回
		if mask[int(b-'a')] > 1 {
			return
		}
	}
	for i := index + 1; i <= len(arr); i++ {
		helper(i, arr, mask, count+temp)
	}

}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
/*
给定一个Excel表格中的列名称，返回其相应的列序号。

例如，

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/excel-sheet-column-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func titleToNumber(columnTitle string) int {
	//相当于26进制
	b := []byte(columnTitle)
	if len(b) == 0 {
		return 0
	}
	ret := 0
	temp := 1
	for i := len(b) - 1; i >= 0; i-- {
		ret += int(b[i]-'A'+1) * temp
		temp *= 26
	}
	return ret
}
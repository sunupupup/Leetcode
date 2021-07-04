func findErrorNums(nums []int) []int {
	mask := make([]int, len(nums))
	a, b := 0, 0

	for _, v := range nums {
		mask[v-1]++ //进行统计，肯定有一个数字出现 0次，一个数字出现2次 ，就是个最简单的哈希
	}

	for i, v := range mask {
		if a != 0 && b != 0 { //找到了的话提前结束遍历
			break
		}
		if v == 0 { //丢掉的数
			a = i + 1
		} else if v == 2 { //重复的数
			b = i + 1
		}
	}
	return []int{b, a}
}

//官方：位运算
func findErrorNums(nums []int) []int {
	//位运算
	//推理：失去的数字和重复的数字，出现的次数奇偶性一样
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	n := len(nums)
	for i := 1; i <= len(nums); i++ {
		xor ^= i
	}
	//此时xor相当于  缺失的数 ^ 重复的数
	lowbit := xor & -xor
	num1, num2 := 0, 0
	for _, v := range nums {
		if v&lowbit == 0 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}
	for i := 1; i <= n; i++ {
		if i&lowbit == 0 {
			num1 ^= i
		} else {
			num2 ^= i
		}
	}
	for _, v := range nums {
		if v == num1 {
			return []int{num1, num2}
		}
	}
	return []int{num2, num1}

}

//数学方法：
// sum(nums) - sum(set(nums)) = 重复的数字		set就是去重的操作
// (1 + len(nums)) * len(nums) - sum(set(nums)) = 丢失的数字
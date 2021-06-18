
//数学方法 利用 k^m < n=k^0+k^1+...+k^m  < (k+1)^m
func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)
	for i := 63; i > 1; i-- {

		k := int64(math.Pow(float64(num), 1.0/float64(i)))
		if k == 1 {
			continue
		}
		var temp int64
		var kk int64 = 1
		for j := 0; j <= i; j++ {
			temp += kk
			kk *= k
		}
		if temp == int64(num) {
			return strconv.Itoa(int(k))
		}

	}
	return strconv.Itoa(num - 1)
}

//暴力法，只适用于不超过int的数，也不能太大
//"1000000000000000000" 这中直接无法转为int，且毕然超时
func smallestGoodBase(n string) string {
	//n进制  n^0 + n^1 + n^2 +++ n^m = num
	//n 属于 [2,n-1]   很显然 n-1 进制必然是一个答案
	//暴力法试一下
	num, _ := strconv.Atoi(n)
	for i := 2; i < num; i++ {
		sum := 1
		temp := 1
		for sum < num {
			sum += temp * i
			temp *= i
			if sum == num {
				return strconv.Itoa(i)
			}
		}
	}
	return ""
}
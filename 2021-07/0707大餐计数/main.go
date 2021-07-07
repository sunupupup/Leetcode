//我的方法，先hash统计出所有，然后 n^2 复杂度，每两个计算一次结果
func countPairs(deliciousness []int) int {

	foodsCount := make(map[int]int, 0)
	for i := range deliciousness {
		//if _,ok := foodsCount[deliciousness[i]];!ok {
		//    foodsCount[deliciousness[i]] = 1
		//}else{
		//    foodsCount[deliciousness[i]] += 1
		//}
		foodsCount[deliciousness[i]] += 1
	}

	var CombinationBetweenTwoFood = func(a, b int) int {
		if a == b {
			return foodsCount[a] * (foodsCount[a] - 1) / 2
		}
		return foodsCount[a] * foodsCount[b]
	}

	deliciousnessKind := []int{}
	for k, _ := range foodsCount {
		deliciousnessKind = append(deliciousnessKind, k)
	}

	fmt.Println(deliciousnessKind)

	ret := 0
	for i := 0; i < len(deliciousnessKind); i++ {
		deliciousness1 := deliciousnessKind[i]
		for j := i; j < len(deliciousnessKind); j++ {
			deliciousness2 := deliciousnessKind[j]
			if isMatch(deliciousness1 + deliciousness2) {
				fmt.Println(deliciousness1, " ", deliciousness2)
				ret += CombinationBetweenTwoFood(deliciousness1, deliciousness2)
				fmt.Println(ret)

			}
		}
	}
	return ret % (1000000000 + 7)
}

//判断一个数是不是2的整次幂
func isMatch(number int) bool {
	if (number&(number-1)) == 0 && number != 0 { //0不是2的整次幂，要排除
		return true
	}
	return false
}

//官方的思路：先找到最大的，无需去重，然后一边遍历一边放到map里面
func countPairs2(deliciousness []int) (ret int) {

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxDeliciousness := deliciousness[0]
	for i := range deliciousness {
		maxDeliciousness = max(maxDeliciousness, deliciousness[i])
	}
	maxSum := 2 * maxDeliciousness
	cnt := map[int]int{}
	for _, val := range deliciousness {
		for sum := 1; sum <= maxSum; sum = sum << 1 {
			ret += cnt[sum-val]
		}
		cnt[val]++
	}
	return ret % (1000000000 + 7)
}


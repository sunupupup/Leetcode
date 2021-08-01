/*
贪心 + 排序

以下推导摘自leetcode评论

贪心唯一的难点就是证明。通常数学直觉都能告诉我们：“最大”与“最小”必须匹配成一对，才能使得各项数对和尽量“均匀”，才能使最大和达到最小。

暂避纯粹的数学公式，给一点思路。我们如果要证明这题，还是基于反证假设——假设“最大”与“最小”不成对，能得到更小的数对和，这里分情况讨论：

1. 想拆散(min, max)，可以找一个其他数对中的数来替换min，假设拆散(x, y)，这里有min <= x <= y <= max：
    1.1 我们拿较小的x来替换min，得到(x, max)和(min, y)：
        由于x + max >= min + max，x + max >= x + y，又x + max >= min + y，
        所以新得到的数对(x, max)将成为不小于原先组合的值，则不能得到更小的最大数对和，与假设矛盾；
    1.2 我们拿较小的y来替换min，同1.1，新数对(y, max)将比1.1中的更大，则不能得到更小的数对和；
2. 想拆散(min, max)，可以找一个其他数对中的数来替换max，不难得到和1.中对偶的结论。
如果(min, max)必须为一对，那么你很容易证明“次大”与“次小”必须为一对——它既不能和(min, max)做对换，也不可能和他们俩之间的数对进行对换，依此类推——“排序后的第k个和倒数第k个元素必须成对”得证。
*/

func minPairSum(nums []int) int {
	//[max + min] 到 [max1 + max2]
	sort.Ints(nums)
	n := len(nums)
	ret := 0
	for i := 0; i < n/2; i++ {
		ret = max(ret, nums[i]+nums[n-1-i])
	}
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
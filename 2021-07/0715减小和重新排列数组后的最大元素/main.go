/*
给你一个正整数数组 arr 。请你对 arr 执行一些操作（也可以不进行任何操作），使得数组满足以下条件：

arr 中 第一个 元素必须为 1 。
任意相邻两个元素的差的绝对值 小于等于 1 ，也就是说，对于任意的 1 <= i < arr.length （数组下标从 0 开始），都满足 abs(arr[i] - arr[i - 1]) <= 1 。abs(x) 为 x 的绝对值。
你可以执行以下 2 种操作任意次：

减小 arr 中任意元素的值，使其变为一个 更小的正整数 。
重新排列 arr 中的元素，你可以以任意顺序重新排列。
请你返回执行以上操作后，在满足前文所述的条件下，arr 中可能的 最大值 。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-element-after-decreasing-and-rearranging
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//这个方法时间复杂度是 O(n lgn)
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    sort.Ints(arr)
    arr[0] = 1
    for i:=1;i<len(arr);i++ {
        if arr[i]-arr[i-1] > 1 {
            arr[i] = arr[i-1] + 1
        }
    }
    return arr[len(arr)-1]

}

//考虑 O(n)的方法，

/*
深挖题目隐含的性质，我们可以将时间复杂度优化至O(n)。
记arr 的长度为 n。由于第一个元素必须为 1，且任意两个相邻元素的差的绝对值小于等于 1，故答案不会超过 n。
所以我们只需要对arr 中值不超过 n 的元素进行计数排序，而对于值超过 n 的元素，由于其至少要减小到 n，我们可以将其视作 n。
读者可据此修改方法一中的排序代码，此处不再赘述，我们将重点转到另一种计算策略上。
从另一个视角来看，为了尽可能地构造出最大的答案，我们相当于是在用 arr 中的元素去填补自身在[1,n] 中缺失的元素。
首先，我们用一个长为n+1 的数组 cnt 统计arr 中的元素个数（将值超过 n 的元素视作 n）。然后，从 1 到 n 遍历 cnt 数组，
若 cnt[i]=0，则说明缺失元素 ii，我们需要在后续找一个大于 ii 的元素，将其变更为 i。
我们可以用一个变量miss 记录cnt[i]=0 的出现次数，当遇到cnt[i]>0 时，则可以将多余的cnt[i]−1 个元素减小，补充到之前缺失的元素上。
遍历cnt 结束后，若此时miss=0，则说明修改后的arr 包含了[1,n] 内的所有整数；否则，对于不同大小的缺失元素，我们总是优先填补较小的，因此剩余缺失元素必然是[n−miss+1,n] 这一范围内的miss 个数，因此答案为 n−miss。

*/
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    n := len(arr)
    cnt := make([]int, n+1)
    for _, v := range arr {
        cnt[min(v, n)]++
    }
    miss := 0
    for _, c := range cnt[1:] {
        if c == 0 {
            miss++
        } else {
            miss -= min(c-1, miss) // miss 不会小于 0，故至多减去 miss 个元素
        }
    }
    return n - miss
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
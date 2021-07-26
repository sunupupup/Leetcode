


//自己的做法  就遍历所要求的区间所有数字
func isCovered(ranges [][]int, left int, right int) bool {
    //最简单就是一个个数字遍历过去，反正是有限个数
    for i:=left;i<=right;i++ {
        is_covered := false
        for _,v := range ranges {
            if i >= v[0] && i <= v[1] {
                is_covered = true
                break
            }
        }
        if is_covered == false {
            return false
        }
    }
    return true
}

//官方：差分数组 
//常规思路：遍历每个区间，对结果区间每个数+1，遍历完所有区间之后，查看给定区间内所有值是否大于1，假设ranges长度为n，结果区间长度l，这个时间复杂度是 l * n
//利用差分数组，就是 l + n 的复杂度
//用差分数组 diff 维护相邻两个整数的被覆盖区间数量变化量




//找到所有的1的位置,以goal为间隔，然后搜索符合要求的两个1的区间左右有多少0
func numSubarraysWithSum(nums []int, goal int) int {
    //先找到所有1的位置
    s := []int{}    //存放所有1的index
    for i:=0;i<len(nums);i++ {
        if nums[i]==1 {
            s = append(s,i)
        }
    }
    if len(s) < goal {
        return 0
    }
    if len(s)==0 {
        return (1+len(nums))*(len(nums))/2
    }
    //找到所有1，左右的0
    m := make(map[int][2]int,0)
    for i:=0;i<len(s);i++ {
        if i!=0 {
            index2 := s[i]+1
            r := 0
            for ;index2<len(nums) && nums[index2]!=1;index2++ {
                r++
            }
            m[s[i]] = [2]int{m[s[i-1]][1],r}
        }else{
            index1,index2 := s[i]-1,s[i]+1
            l,r := 0,0
            for ;index1>=0 && nums[index1]!=1;index1-- {
                l++
            }
            for ;index2<len(nums) && nums[index2]!=1;index2++ {
                r++
            }
            m[s[i]] = [2]int{l,r}
        }
    }
    ret := 0
    if goal==0 {
        //对所有长度的0进行
        ret += (m[s[0]][0]+1)*m[s[0]][0]/2
        for i:=0;i<len(s);i++{
            ret += (m[s[i]][1]+1)*m[s[i]][1]/2
        }
    }else{
        for i:=0;i<=len(s)-goal;i++{
            ret += (m[s[i]][0]+1)*(m[s[i+goal-1]][1]+1)
        }
    }

    return ret

}

//官方解法，看完之后发现自己写的就是垃圾。。。
//sum[j] - sum[i] = goal
func numSubarraysWithSum2(nums []int, goal int) int {
    //前缀和
    m := make(map[int]int,0)
    ret :=0
    sum := 0
    for _,num := range nums{
        m[sum]++
        sum += num
        ret += m[sum - goal]
    }
    return ret
}

//滑动窗口
func numSubarraysWithSum3(nums []int, goal int) (ans int) {
    left1, left2 := 0, 0
    sum1, sum2 := 0, 0
    for right, num := range nums {
        sum1 += num
        for left1 <= right && sum1 > goal {
            sum1 -= nums[left1]
            left1++
        }
        sum2 += num
        for left2 <= right && sum2 >= goal {
            sum2 -= nums[left2]
            left2++
        }
        ans += left2 - left1
    }
    return
}

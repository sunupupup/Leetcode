//

func numberOfArithmeticSlices(nums []int) int {
    //子数组：连续序列
    n := len(nums)
    now := 0
    ret := 0
    for now<n {
        //从now出发，向后找
        interval := 0
        if now<n-1 {
            interval = nums[now+1]-nums[now]
        }else{
            break
        }

        temp := now
        for now<n-1 {
            if nums[now]+interval==nums[now+1] {	//向后找等差的序列
                now++
            }else{
                break
            }
        }
        ret += helper(now-temp+1)
    }
    return ret
}

func helper(a int)int{
    if a <3 {
        return 0
    }
    return (a-1)*(a-2)/2
}
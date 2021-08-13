/*
给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
示例 1：

输入：n = 13
输出：6
*/





//暴力法
func countDigitOne(n int) int {
    ret := 0
    for i:=1;i<=n;i++ {
        if i%10==0 {
            continue
        }
        count := countOne(i)
        if count!=0 {
            for num:=i;num <= n;num*=10 {
                ret += count
            }
        }

    }
    return ret
}

func countOne(a int)int{
    ret := 0
    for a>0 {
        if a%10 == 1{
            ret++
        }
        a /= 10
    }
    return ret
}
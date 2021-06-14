package main 

type ListNode struct {
    Val int
    Next *ListNode
}
 
//法一 ： 足够大的数组
 func reversePrint1(head *ListNode) []int {
    ret := make([]int,10000,10000)
    index :=9999
    temp := head
    for temp!=nil{
        ret[index]= temp.Val
        index--
        temp = temp.Next
    }
    return ret[index+1:]
}

//直接获取 反转
func reversePrint2(head *ListNode) []int {
    ret := make([]int,0)
    temp := head
    for temp!=nil {
        ret = append(ret,temp.Val)
        temp=temp.Next
    }
    n := len(ret)
    for i:=0;i<n/2;i++{
        ret[i],ret[n-1-i] = ret[n-1-i],ret[i]		//反转一个数组
    }
    return ret[:n]
}
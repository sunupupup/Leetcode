package main
import . "nc_tools"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 
 * @param m int整型 
 * @param a ListNode类 指向彩带的起点，val表示当前节点的val，next指向下一个节点
 * @return ListNode类一维数组
*/
func solve( m int ,  a *ListNode ) []*ListNode {
    // write code here
    temp := make([][]int,m)
    
    for a != nil {
        index := a.Val % m
        temp[index] = append(temp[index],a.Val)
        a = a.Next
    }
    
    ret := []*ListNode{}
    
    for i:=0;i<m;i++ {
        ret = append(ret,createList(temp[i]))
    }
    
    return ret
    
}

//创建链表，返回头节点
func createList(a []int)*ListNode{
    if len(a) == 0 {
        return nil
    }
    head := &ListNode{a[0],nil}
    temp := head
    for i:=1;i<len(a);i++ {
        b := &ListNode{a[i],nil}
        temp.Next = b
        temp = b
    }
    return head
    
}
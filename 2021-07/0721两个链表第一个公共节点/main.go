/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 //双指针法
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    headA_temp := headA
    headB_temp := headB
    count := 0
    for headA_temp !=headB_temp {
        if headA_temp.Next != nil {
            headA_temp = headA_temp.Next
        }else{
            count++
            if count > 1 {
                return nil
            }
            headA_temp = headB
        }

        if headB_temp.Next != nil {
            headB_temp = headB_temp.Next
        }else{
            headB_temp = headA
        } 
    }
    return headA_temp
}

//哈希表法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    //哈希表法O(n)复杂度
    m := map[*ListNode]struct{}{}
    temp := headA
    for temp != nil {
        m[temp] = struct{}{}
        temp = temp.Next
    }
    temp = headB
    for temp != nil {
        if _,ok := m[temp];ok {
            return temp
        }
        temp = temp.Next
    }
    
    return nil

}
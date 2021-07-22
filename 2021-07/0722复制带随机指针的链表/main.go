/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

//用map记下位置
//直接可利用数组构建链表
func copyRandomList(head *Node) *Node {
	//用一个数组先记录下来，然后每个数组的值，是randonm的位置
	n := 0
	m := make(map[*Node]int, 0) //每个节点对应的下标
	if head == nil {
		return nil
	}
	temp := head
	for temp != nil {
		m[temp] = n
		n++
		temp = temp.Next
	}

	type Pair [2]int //记录当前值，和random在链表中的索引位置
	arr := make([]Pair, n, n)
	temp = head
	for i := range arr {
		index := -1
		if temp.Random != nil {
			index = m[temp.Random]
		}
		arr[i] = Pair{temp.Val, index}
		temp = temp.Next
	}

	//fmt.Println(arr)

	//根据arr来构建链表
	ret := make([]Node, n)
	for i := 0; i < n; i++ {
		if arr[i][1] == -1 {
			ret[i] = Node{Val: arr[i][0]}
		} else {
			ret[i] = Node{Val: arr[i][0], Random: &ret[arr[i][1]]}
		}
		if i < n-1 {
			ret[i].Next = &ret[i+1]
		}
	}

	return &ret[0]

}

//官方
var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next) //递归，深拷贝
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}
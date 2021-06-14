package main

//TreeNode ...Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(v []*TreeNode) (ret []*TreeNode) {
	for _, i := range v {
		if i.Left != nil {
			ret = append(ret, i.Left)
		}
		if i.Right != nil {
			ret = append(ret, i.Right)
		}
	}
	return
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	left2right := true
	//存放每一层的节点
	v := []*TreeNode{root}
	var ret [][]int
	if root == nil {
		return ret
	}
	for len(v) != 0 {
		n := len(v)
		if left2right {
			var retpart []int
			for i := 0; i < n; i++ {
				retpart = append(retpart, v[i].Val)
			}
			v = dfs(v)
			ret = append(ret, retpart)
			left2right = false
		} else { //从右到左

			var retpart []int
			for i := n - 1; i >= 0; i-- {
				retpart = append(retpart, v[i].Val)
			}
			v = dfs(v)
			ret = append(ret, retpart)
			left2right = true
		}
	}

	return ret
}

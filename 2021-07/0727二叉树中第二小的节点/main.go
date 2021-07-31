/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	//宽搜,一边搜一边保存
	min1 := root.Val
	min2 := root.Val

	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		temp := []*TreeNode{}
		for _, v := range nodes {
			if min1 == min2 && v.Val > min2 {
				min2 = v.Val
			} else if v.Val > min1 && v.Val < min2 {
				min2 = v.Val
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
				temp = append(temp, v.Left)
			}
		}
		nodes = temp

	}
	if min1 == min2 {
		return -1
	}
	return min2
}

//深搜
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	//深搜
	min1 := root.Val
	min2 := root.Val

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if (min1 == min2 && node.Val > min2) || (min1 != min2 && node.Val < min2 && node.Val > min1) {
			min2 = node.Val
		}
		if node.Left != nil {
			dfs(node.Left)
			dfs(node.Right)
		}

	}
	dfs(root)
	if min1 == min2 {
		return -1
	}
	return min2
}
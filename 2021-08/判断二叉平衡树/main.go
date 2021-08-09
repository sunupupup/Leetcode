
//判断是不是二叉平衡树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//这是自底向上的做法
func isBalanced(root *TreeNode) bool {

	var getHeight func(*TreeNode) (int, bool)
	getHeight = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}

		l, b1 := getHeight(node.Left)
		r, b2 := getHeight(node.Right)
		if b1 && b2 {
			if l-r <= 1 && l-r >= -1 {
				return max(l, r) + 1, true
			}
		}
		return 0, false
	}

	_, flag := getHeight(root)

	return flag

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

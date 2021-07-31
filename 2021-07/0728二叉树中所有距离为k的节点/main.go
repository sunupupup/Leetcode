/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//哈希表 + 深搜
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	//向下很容易，直接从target往下搜就行了
	//向上的话,需要记录父节点
	//哈希表记录每个节点的父节点
	if k == 0 {
		return []int{target.Val}
	}
	//先记录所有节点的父亲
	m := map[int]*TreeNode{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		//传进来的是父节点
		if node.Left != nil {
			m[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			m[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)

	//从target出发  dfs  寻找所有深度为k的节点
	//一层一层遍历
	ret := []int{}
	var findAns func(*TreeNode, *TreeNode, int)
	//from节点记录从哪来的，防止回头  ！！！ 这里很重要
	findAns = func(node, from *TreeNode, depth int) {

		if node == nil {
			return
		}
		if depth == k {
			ret = append(ret, node.Val)
			return
		}
		//判断左子节点、右子节点、父节点，并判断是否回头
		if m[node.Val] != from {
			findAns(m[node.Val], node, depth+1)
		}
		if node.Left != from {
			findAns(node.Left, node, depth+1)
		}
		if node.Right != from {
			findAns(node.Right, node, depth+1)
		}
	}
	findAns(target, nil, 0)
	return ret

}







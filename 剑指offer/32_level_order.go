package offer

func levelOrder(root *TreeNode) []int {
	var result []int
	tree := []*TreeNode{root}
	for tree != nil {
		var hierarchyTree []*TreeNode
		for i := range tree {
			result = append(result, tree[i].Val)
			if tree[i].Left != nil {
				hierarchyTree = append(hierarchyTree, tree[i].Left)
			}
			if tree[i].Right != nil {
				hierarchyTree = append(hierarchyTree, tree[i].Right)
			}
		}
		tree = hierarchyTree
	}
	return result
}

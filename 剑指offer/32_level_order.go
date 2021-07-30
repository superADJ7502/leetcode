package offer

// 1
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

// 2
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	tree := []*TreeNode{root}
	for tree != nil {
		var childResult []int
		var hierarchyTree []*TreeNode
		for i := range tree {
			if tree[i].Left != nil {
				hierarchyTree = append(hierarchyTree, tree[i].Left)
			}
			if tree[i].Right != nil {
				hierarchyTree = append(hierarchyTree, tree[i].Right)
			}
			childResult = append(childResult, tree[i].Val)
		}
		result = append(result, childResult)
		tree = hierarchyTree
	}
	return result
}

// 3
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var reverseMark bool
	tree := []*TreeNode{root}
	for tree != nil {
		var childResult []int
		var hierarchyTree []*TreeNode
		for i := range tree {
			if tree[i].Left != nil {
				hierarchyTree = append(hierarchyTree, tree[i].Left)
			}
			if tree[i].Right != nil {
				hierarchyTree = append(hierarchyTree, tree[i].Right)
			}
			childResult = append(childResult, tree[i].Val)
		}
		if reverseMark {
			reverse(childResult)
		}
		result = append(result, childResult)
		tree = hierarchyTree
		reverseMark = !reverseMark
	}
	return result
}

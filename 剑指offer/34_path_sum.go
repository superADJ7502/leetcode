package offer

func pathSum(root *TreeNode, target int) [][]int {
	var result [][]int
	pathSumHelper(root, target, nil, &result)
	return result
}

func pathSumHelper(root *TreeNode, target int, values []int, result *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root.Val == target {
		final := make([]int, len(values))
		copy(final, values)
		*result = append(*result, append(final, target))
		return
	}
	pathSumHelper(root.Left, target-root.Val, append(values, root.Val), result)
	pathSumHelper(root.Right, target-root.Val, append(values, root.Val), result)
}

package util

// TreeNode 节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SliceToTree
// 1. list[0]无效
// 2. idx start from 1
func SliceToTree(list []int, idx int) *TreeNode {
	if idx >= len(list) {
		return nil
	}
	return &TreeNode{
		Val:   list[idx],
		Left:  SliceToTree(list, 2*idx),
		Right: SliceToTree(list, 2*idx+1),
	}
}

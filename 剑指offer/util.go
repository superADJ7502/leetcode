package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// TreeNode 节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

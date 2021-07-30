package offer

// ----------------- List -----------------

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// ----------------- tree -----------------

// TreeNode 节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SliceToTree
// default idx==1
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

// ----------------- compare -----------------
func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

// ----------------- reverse -----------------

func reverse(values []int) {
	l := len(values)
	for i := 0; i < l/2; i++ {
		values[i], values[l-1-i] = values[l-1-i], values[i]
	}
}

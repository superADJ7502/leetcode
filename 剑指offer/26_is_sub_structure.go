package offer

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	return isSubStructureHelper(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSubStructureHelper(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return isSubStructureHelper(A.Left, B.Left) && isSubStructureHelper(A.Right, B.Right)
}

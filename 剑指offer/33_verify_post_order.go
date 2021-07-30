package offer

func verifyPostorderHelper(postorder []int, left, right int) bool {
	if left >= right {
		return true
	}
	mid := left
	for postorder[mid] < postorder[right] {
		mid++
	}
	for postorder[mid] > postorder[right] {
		mid++
	}
	return postorder[mid] == postorder[right] && verifyPostorderHelper(postorder, left, mid-1) && verifyPostorderHelper(postorder, mid, right-1)
}

func verifyPostorder(postorder []int) bool {
	return verifyPostorderHelper(postorder, 0, len(postorder)-1)
}

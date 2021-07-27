package offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	start, end := head, head
	for k != 0 {
		end = end.Next
		k--
	}
	for end != nil {
		start, end = start.Next, end.Next
	}
	return start
}

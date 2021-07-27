package offer

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

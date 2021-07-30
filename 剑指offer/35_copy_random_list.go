package offer

func copyRandomList(head *Node) *Node {
	hash := make(map[*Node]*Node)
	l, newList := head, &Node{}
	ans, newHead := newList, newList
	for l != nil {
		node := &Node{Val: l.Val}
		newList.Next = node
		hash[l] = node
		newList = newList.Next
		l = l.Next
	}
	newHead = newHead.Next
	for head != nil {
		newHead.Random = hash[head.Random]
		newHead = newHead.Next
		head = head.Next
	}
	return ans.Next
}

package piscine

func ListReverse(l *List) {
	var current, next, prev *NodeL
	if l.Head == nil {
		return
	}
	l.Tail = l.Head
	current = l.Head
	next = current.Next
	for next != nil {
		current.Next = prev
		prev = current
		current = next
		next = next.Next
	}
	current.Next = prev
	l.Head = current
}

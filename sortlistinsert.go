package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	var prv *NodeI
	node := &NodeI{data_ref, nil}
	curr := l
	if curr == nil {
		l = node
		return l
	}
	for {
		if curr == nil {
			prv.Next = node
			break
		}
		if curr.Data > data_ref {
			if prv != nil {
				node.Next = curr
				prv.Next = node
			} else {
				node.Next = l
				l = node
			}
			break
		}
		prv = curr
		curr = curr.Next
	}
	return l
}

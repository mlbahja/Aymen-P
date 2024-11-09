package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	change := true
	for change {
		change = false
		current := l
		for current != nil && current.Next != nil {
			if current.Data > current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				change = true
			}
			current = current.Next
		}
	}
	return l
}

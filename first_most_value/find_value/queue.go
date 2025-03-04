package findvalue

type queue []*Nodes

func NewQueueWithSize() queue {
	return make([]*Nodes, 0)
}

func (q *queue) Push(value *Nodes) {
	*q = append(*q, value)
}

func (q *queue) Pop() *Nodes {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) IsFull(queueSize int) bool {
	return len(*q) == queueSize
}

func (q *queue) GetByIndex(index int) *Nodes {
	return (*q)[index]
}

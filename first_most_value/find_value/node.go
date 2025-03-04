package findvalue

type Nodes struct {
	Nodes   map[string]Summary
	MaxNode *Summary
}

func NewNodes() Nodes {
	maxNode := Summary{previousNode: "", Sum: 0}
	return Nodes{
		Nodes:   make(map[string]Summary),
		MaxNode: &maxNode,
	}
}

func (m Nodes) Add(key string, value Summary) {
	if _, ok := m.Nodes[key]; !ok {
		m.Nodes[key] = value
	} else {
		if m.Nodes[key].Sum < value.Sum {
			m.Nodes[key] = value
		}
	}
	if m.MaxNode.Sum < value.Sum {
		m.MaxNode.previousNode = key
		m.MaxNode.Sum = value.Sum
	}
}

func (m Nodes) Get() *Nodes {
	return &m
}

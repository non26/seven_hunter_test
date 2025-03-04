package findvalue

import (
	"fmt"
)

type Summary struct {
	previousNode string
	Sum          int
}

func CreateKey(level int, index int) string {
	return fmt.Sprintf("%v_%v", level, index)
}

func FindValue(input [][]int) int {
	queueSize := 2
	queue := NewQueueWithSize()
	startNodes := NewNodes()
	startNodes.Add("0_0", Summary{previousNode: "", Sum: input[0][0]})
	queue.Push(&startNodes)
	if len(input) == 1 {
		return queue.Pop().MaxNode.Sum
	}
	for level, row := range input[0 : len(input)-1] {
		if queue.IsFull(queueSize) {
			queue.Pop()
		}
		nodes := NewNodes()
		fmt.Printf("level: %v\n", level)
		for currentNodeIndex := range row {
			currentKey := CreateKey(level, currentNodeIndex)

			nextLevel := level + 1
			neighborOfCurrentNode := []int{currentNodeIndex, currentNodeIndex + 1}

			key_nb1 := CreateKey(nextLevel, neighborOfCurrentNode[0])
			sum_nb1 := queue.GetByIndex(0).Nodes[currentKey].Sum + input[nextLevel][neighborOfCurrentNode[0]]
			nodes.Add(key_nb1, Summary{previousNode: currentKey, Sum: sum_nb1})

			key_nb2 := CreateKey(nextLevel, neighborOfCurrentNode[1])
			sum_nb2 := queue.GetByIndex(0).Nodes[currentKey].Sum + input[nextLevel][neighborOfCurrentNode[1]]
			nodes.Add(key_nb2, Summary{previousNode: currentKey, Sum: sum_nb2})
		}
		queue.Push(nodes.Get())
	}

	fmt.Printf("Max Value: %v\n", queue.GetByIndex(queueSize-1).MaxNode.Sum)

	return queue.GetByIndex(queueSize - 1).MaxNode.Sum
}

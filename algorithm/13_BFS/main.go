package main

import "fmt"

func (n *GraphNode) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type GraphNode struct {
	Value         int
	Visited       bool
	AdjacencyList []*GraphNode
}

type Queue struct {
	nodes []*GraphNode
}

func (q *Queue) String() string {
	s := ""
	for _, n := range q.nodes {
		s += fmt.Sprintf("%v ", n)
	}
	return s
}
func NewQueue() *Queue {
	return &Queue{nodes: []*GraphNode{}}
}

func (q *Queue) enqueue(n *GraphNode) {
	q.nodes = append(q.nodes, n)
}

func (q *Queue) dequeue() *GraphNode {
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

func main() {

	// Graph
	// 1-> 2,3
	// 2-> 1,3,4,5
	// 3-> 1,2,6 ,7
	// 4-> 2,5
	// 5-> 2,4
	// 6-> 3,7
	// 7-> 3,6
	node1 := &GraphNode{Value: 1}
	node2 := &GraphNode{Value: 2}
	node3 := &GraphNode{Value: 3}
	node4 := &GraphNode{Value: 4}
	node5 := &GraphNode{Value: 5}
	node6 := &GraphNode{Value: 6}
	node7 := &GraphNode{Value: 7}

	node1.AdjacencyList = append(node1.AdjacencyList, node2)
	node1.AdjacencyList = append(node1.AdjacencyList, node3)

	node2.AdjacencyList = append(node2.AdjacencyList, node1)
	node2.AdjacencyList = append(node2.AdjacencyList, node3)
	node2.AdjacencyList = append(node2.AdjacencyList, node4)
	node2.AdjacencyList = append(node2.AdjacencyList, node5)

	node3.AdjacencyList = append(node3.AdjacencyList, node1)
	node3.AdjacencyList = append(node3.AdjacencyList, node2)
	node3.AdjacencyList = append(node3.AdjacencyList, node6)
	node3.AdjacencyList = append(node3.AdjacencyList, node7)

	node4.AdjacencyList = append(node4.AdjacencyList, node2)
	node4.AdjacencyList = append(node4.AdjacencyList, node5)

	node5.AdjacencyList = append(node5.AdjacencyList, node2)
	node5.AdjacencyList = append(node5.AdjacencyList, node4)

	node6.AdjacencyList = append(node6.AdjacencyList, node3)
	node6.AdjacencyList = append(node6.AdjacencyList, node7)

	node7.AdjacencyList = append(node7.AdjacencyList, node3)
	node7.AdjacencyList = append(node7.AdjacencyList, node6)

	// BFS
	BFS(node1)

}

func BFS(node *GraphNode) {
	queue := NewQueue()
	node.Visited = true
	queue.enqueue(node)

	for len(queue.nodes) > 0 {
		element := queue.dequeue()
		fmt.Println(element.Value)
		for _, n := range element.AdjacencyList {
			if !n.Visited {
				n.Visited = true
				queue.enqueue(n)
			}
		}
	}
}

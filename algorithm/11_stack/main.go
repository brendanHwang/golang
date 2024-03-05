package main

import "fmt"

type Node struct {
	Value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	nodes []*Node
	count int
}

func (s *Stack) String() string {
	result := ""

	result += "nodes: "
	for _, node := range s.nodes[:s.count] {
		result += node.String() + ", "
	}

	result += fmt.Sprintf("\ncount: %v", s.count)

	return result

}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func main() {
	s := NewStack()

	s.Push(&Node{3})
	s.Push(&Node{2})
	s.Push(&Node{1})

	s.Pop()

	fmt.Println(s)
}

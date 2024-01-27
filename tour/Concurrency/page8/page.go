package page8

import (
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
	close(ch)
}

func EquivalentBinaryTrees() {
	//ch := make(chan int, 20)
	//go Walk(tree.New(1), ch)
	//for i := range ch {
	//	println(i)
	//}
	println(Same(tree.New(1), tree.New(1))) // true
	println(Same(tree.New(1), tree.New(2))) // false

}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

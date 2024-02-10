package page13

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func NillInterfaceValues() {
	var i I
	describe(i)
	i.M()
}

package page2

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Insert(val T) {
	l.next = &List[T]{next: l.next, val: val}
}

func (l *List[T]) String() string {
	var builder strings.Builder
	for l != nil {
		fmt.Fprintf(&builder, "%v -> ", l.val)
		l = l.next
	}
	return builder.String()
}

func GenericTypes() {
	l := List[int]{
		val: 1,
	}

	l.Insert(3)
	l.Insert(4)

	fmt.Println(l.String())

}

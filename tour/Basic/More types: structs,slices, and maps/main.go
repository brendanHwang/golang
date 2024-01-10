package main

import (
	"fmt"
	"strings"
)

// page1
func Pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer [예상 결과 42]
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i [예상 결과 21]

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j [예상 결과 2701/37]
}

// page2
type Vertex struct {
	X int
	Y int
}

func Structs() {
	fmt.Println(Vertex{1, 2})
}

// page3

func StructFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

// page4
func PointersToStructs() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

// page5
type Vertex2 struct {
	X, Y int
}

func StructLiterals() {
	v1 := Vertex2{1, 2}
	v2 := Vertex2{X: 1}
	v3 := Vertex2{}
	p := &Vertex2{1, 2}
	fmt.Println(v1, p, v2, v3)

	v4 := Vertex{
		Y: 1,
	}
	fmt.Println(v4)
}

// page6

func Arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// page7
func Slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	array := [6]int{2, 3, 5, 7, 11, 13}
	a := array[1:4]
	fmt.Println(a)

}

// page8
func SliceAreLinkReferenceToArray() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// page9
func SliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// page10

func SliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

// page11

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}

func SliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

// page12
func NilSlices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

// page13
func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func CreatingASliceWithMake() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)

	e := d[2:3]
	printSlice2("e", e)

	f := e[1:]
	printSlice2("f", f)
}

// page14
func SlicesOfSlices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func main() {
	//Pointers()
	//Structs()
	//StructFields()
	//PointersToStructs()
	//StructLiterals()
	//Arrays()
	//Slices()
	//SliceAreLinkReferenceToArray()
	//SliceLiterals()
	//SliceDefaults()
	//SliceLengthAndCapacity()
	//NilSlices()
	//CreatingASliceWithMake()
	SlicesOfSlices()
}

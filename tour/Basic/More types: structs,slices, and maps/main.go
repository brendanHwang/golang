package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"math"
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

// page15
func AppendingToASlice() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

// page16
func Range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// page17

func RangeContinued() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}

// page18
func Pic(dx, dy int) [][]uint8 {
	// 2차원 슬라이스 생성
	image := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		// 각 행에 대한 슬라이스 할당
		image[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// 각 픽셀 값을 계산 (예: (x + y) / 2)
			image[y][x] = uint8((x + y) / 2)
		}
	}
	return image
}

func ExerciseSlices() {
	pic.Show(Pic)
}

// page19

type Vertex3 struct {
	Lat, Long float64
}

var m map[string]Vertex3

func Maps() {
	m = make(map[string]Vertex3)
	m["Bell Labs"] = Vertex3{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

// page20

var m2 = map[string]Vertex3{
	"Bell Labs": Vertex3{
		40.68433, -74.39967,
	},
	"Google": Vertex3{
		37.42202, -122.08408,
	},
}

func MapLiterals() {
	fmt.Println(m2)
}

// page21
func MapLiteralsContinued() {
	m3 := map[string]Vertex3{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)
}

// page22
func MutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// page23
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		m[word]++
	}

	return m
}

func ExerciseMaps() {
	wc.Test(WordCount)
}

// page24

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func FunctionValues() {
	hypot := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}

// page25
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}

}

func FunctionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)

	}
}

// page26
func fibonacci() func() int {
	fib1, fib2 := 0, 1
	return func() int {
		fib1, fib2 = fib2, fib1+fib2
		return fib1
	}

}

func ExerciseFibonacciClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
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
	//SlicesOfSlices()
	//AppendingToASlice()
	//Range()
	//RangeContinued()
	//ExerciseSlices()
	//Maps()
	//MapLiterals()
	//MutatingMaps()
	//ExerciseMaps()
	//FunctionValues()
	//FunctionClosures()
	ExerciseFibonacciClosure()
}

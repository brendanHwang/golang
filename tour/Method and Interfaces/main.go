package main

import (
	"fmt"
	"math"
	"methods/page25"
)

// page1
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// page2
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func MethodsAreFunctions() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

// page3
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func MethodsContinued() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// page4
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func PointerReceivers() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// page5

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func PointerAndFunction() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

// page6
func MethodsAndPointerIndirection() {
	v := Vertex{3, 4}
	v.Scale(2)
	Scale(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	Scale(p, 8)

	fmt.Println(v, p)
}

// page7

func MethodsAndPointerIndirection2() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(Abs(*p))
}

// page8
func ChoosingAValueOrPointerReceiver() {
	v := Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

// page9
type Abser interface {
	Abs() float64
}

func Interface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	a = v

	fmt.Println(a.Abs())
}

// page10
type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func (t T) K() {
	fmt.Println(t.S)

}

func InterfaceAreImplementedImplicitly() {
	var i I = T{"hello"}
	i.M()
}

func main() {
	//Methods()
	//MethodsAreFunctions()
	//MethodsContinued()
	//PointerReceivers()
	//PointerAndFunction()
	//MethodsAndPointerIndirection()
	//MethodsAndPointerIndirection2()
	//ChoosingAValueOrPointerReceiver()
	//Interface()
	//InterfaceAreImplementedImplicitly()
	//page11.InterfaceValues()
	//page12.InterfaceValuesWithNilUnderlyingValues()
	//page13.NillInterfaceValues()
	//page14.TheEmptyInterface()
	//page15.TypeAssertions()
	//page16.TypeSwitches()
	//page17.Stringers()
	//page18.Stringers()
	//page19.Error()
	//page20.Erros()
	//page21.Readers()
	//page22.Readers()
	//page23.Rot13Reader()
	//page24.Image()
	page25.Images()
}

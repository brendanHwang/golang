package page17

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func Stringers() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	var ia, iz fmt.Stringer
	ia, iz = a, z
	fmt.Println(a, z)
	fmt.Println(ia, iz)
}

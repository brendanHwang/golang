package page5

import "fmt"

type car struct {
	make  string
	model string
}

type truck struct {
	car
	bedSize int
}

func EmbededVsNested() {
	// Embeded
	c := car{
		make:  "Toyota",
		model: "Camry",
	}
	t := truck{
		car:     c,
		bedSize: 10,
	}

	fmt.Println(t.make)
	// Nested
	t2 := truck{
		car: car{
			make:  "Toyota",
			model: "Camry",
		},
		bedSize: 10,
	}
	fmt.Println(t2.make)
}

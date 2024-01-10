package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// page1
func For() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Print(sum)
}

// page2

func For_continued() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Print(sum)
}

// page3
func For_is_Go_s_while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// page4
func Forever() {
	for {
	}

}

// page5
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func If() {
	fmt.Println(sqrt(2), sqrt(-4))
}

// page6
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func If_with_a_short_statement() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// page7
func Sqrt(x float64) float64 {
	if x < 0 {
		return 0
	} else if x == 0 {
		return 0
	} else {
		z := 1.0
		for i := 0; i < 10; i++ {
			preZ := z
			z -= (z*z - x) / (2 * z)
			if z == preZ {
				break
			}
		}
		return z
	}
}

func Loop_and_functions() {
	fmt.Println(Sqrt(2))
}

// page9
func Switch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.", os)
	}
}

// page10
func Switch_evaluation_order() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

// page11
func Switch_with_no_condition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

// page12
func Defer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// page13
func Staking_defers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	//For()
	//For_continued()
	//For_is_Go_s_while()
	//Forever()
	//If()
	//If_with_a_short_statement()
	//Loop_and_functions()
	//Switch()
	//Switch_evaluation_order()
	//Switch_with_no_condition()
	//Defer()
	Staking_defers()
}

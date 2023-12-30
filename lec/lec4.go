package lec

import "fmt"

// lec4는 fmt 패키지를 사용하여 출력을 하는 다양한 방법을 보여주는 Go 함수입니다.
//
// 이 함수는 매개변수를 받지 않습니다.
// 반환하는 값이 없습니다.
func lec4() {

	age := 24
	stringAge := "24"
	name := "brendan"

	// Print

	fmt.Print("Hello, ")
	fmt.Print("World!\n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello brendan")
	fmt.Println("goodbye brendan")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings)
	fmt.Printf("my age is %v and my name is %v\n", age, name) // %v 는 default 값을 출력
	fmt.Printf("my age is %d and my name is %s\n", age, name)
	fmt.Printf("my age is %T and my name is %T\n", age, name)
	fmt.Printf("my age is %q and my name is %q\n", stringAge, name)
	fmt.Printf("your scored %f points! \n", 225.555)
	fmt.Printf("your scored %.1f points! \n", 225.555)
	fmt.Printf("my age is %% and my name is %%\n")

	// Sprintf (save formatted strings)
	intro := fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Println(intro)
}

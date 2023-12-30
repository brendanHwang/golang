package main

import (
	"fmt" // formatting
)

func main() {

	// strings
	var nameOne string = "Kim"
	var nameTwo = "Hwnag"
	var nameThree string

	fmt.Print(nameOne, " ", nameTwo, "\n")
	fmt.Printf("%s %s\n", nameOne, nameTwo)
	fmt.Println(nameOne, nameTwo)

	fmt.Print(nameThree, "\n")
	nameThree = "Park"
	fmt.Print(nameThree, "\n")

	nameFour := "Lee"
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	// var numOne int8 = 25 // -128 ~ 127 8bit int
	// var numTwo int16 = 25 // -32768 ~ 32767 16bit int
	// var numThree int32 = 25 // -2147483648 ~ 2147483647 32bit int
	// var numFour int64 = 25 // -9223372036854775808 ~ 9223372036854775807 64bit int

	// float
	var scoreOne float32 = 90.97
	var scoreTwo float64 = 123232314142.3123123123123123123 // 대부분 float64를 사용
	var scoreThree = 90.97                                  // 이렇게해도 float64 사용

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}

// 실행 시키려면 go run lec2.go
// 빌드 하려면 go build lec2.go

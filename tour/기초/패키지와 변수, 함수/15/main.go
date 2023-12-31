// 상수
/*

상수는 변수처럼 선언되지만 const 키워드와 함께 선언됩니다. 특이하게 할당 연산 없이 사용

상수는 character 혹은 string, boolean, 숫자 값이 될 수 있습니다.

상수는 := 를 통해 선언될 수 없습니다.
*/
package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	const World = "世界"

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true

	fmt.Println("Go rules?", Truth)
}

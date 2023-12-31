// 함수 만들기 Go에서는 ( 변수이름 자료형 ) 형태 & 함수의 return에 대한 자료형은 ()옆에 
package main

import(
	"fmt"
)

func add(x int , y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
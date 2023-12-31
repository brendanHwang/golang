/*
Defer
defer문은 자신을 둘러싼 함수가 종료할 때까지 어떠한 함수의 실행을 연기합니다

연기된 호출의 인자는 즉시 평가되지만 그 함수 호출은 자신을 둘러싼 함수가 종료할 때까지 수행되지 않습니다.
*/

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// 두개 반환도 가능!
// 두개 반환  시 (string, string) 이런식으로 반황에 대한 자료형 표시
package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Printf("a: %v, b: %v\n", a, b)
}

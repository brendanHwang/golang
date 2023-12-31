// 기면 반환
// return이라고 쓰고 끝

// 주의 함수 내부에서 := 로 해야 가능 & 반환에 대한 자료형 명시
package main

import (
	"fmt"
)

func nakedReturn(x int) (first, second int) {
	first = x + 1
	second = x * 2
	return
}

func main() {
	a, b := nakedReturn(42)
	fmt.Println(a, b)
}

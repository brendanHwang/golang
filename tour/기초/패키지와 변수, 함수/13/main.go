// Type 변환
/*
C와 달리 Go는 다른 type의 요소들 간의 할당에는 명시적인 변환을 필요로합니다. 예시에서 float64 나 uint 변환을 제거해보시고, 어떤 일이 발생하는 지 보십시오.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)

	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

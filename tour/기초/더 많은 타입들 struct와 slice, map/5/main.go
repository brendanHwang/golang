/*
If
Go의 if 문은 for 반복문과 비슷합니다. 그 표현은 ( ) 괄호로 둘러쌓일 필요는 없지만, { } 괄호는 필수입니다.
*/

package main

import (
	"fmt"
	"math"
	// "math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

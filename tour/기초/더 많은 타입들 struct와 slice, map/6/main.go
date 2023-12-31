/*

If와 else
짧은 if 문 안에서 선언된 변수들은 어떠한 else 블럭에서든 사용이 가능합니다.

pow 에 대한 두 가지 호출 모두 main 의 시작 부분에서 fmt.Println 에 대한 호출하기 이전에 그들의 결과를 반환합니다.

*/


package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

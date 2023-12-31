// 기본 자료형
/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8의 별칭

rune // int32의 별칭
     // 유니코드에서 code point를 의미합니다.

float32 float64

complex64 complex128
*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)

	x1       = 1
	x2 int64 = 2
	x3 int32 = 3
	x4 rune  = 4
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", x1, x1)
	fmt.Printf("Type: %T Value: %v\n", x2, x2)
	fmt.Printf("Type: %T Value: %v\n", x3, x3)
	fmt.Printf("Type: %T Value: %v\n", x4, x4)
}

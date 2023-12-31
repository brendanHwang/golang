/*
For
Go는 for 반복문이라는 단 하나의 반복 구조를 가집니다.

기본적인 for 반복문의 ; 으로 구별되는 세 가지 구성 요소를 갖습니다.

초기화 구문: 첫 번째 iteration 전에 수행됩니다.
조건 표현: 매 모든 iteration 이전에 판별됩니다.
사후 구문: 매 iteration 마지막에 수행됩니다.
초기화 구문은 주로 짧은 변수 선언일 것이며, 거기서 선언된 변수들은 for문의 스코프 내에서만 보여집니다.

반복문은 조건 판별의 boolean 값이 false 이면 iterating을 종료할 것 입니다.

주의: C나 Java, JavaScript와 같은 다른 언어들과 달리, Go는 for문의 세 가지 구성 요소를 감싸는 괄호가 없고, { } 괄호가 항상 필수입니다.
*/
package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

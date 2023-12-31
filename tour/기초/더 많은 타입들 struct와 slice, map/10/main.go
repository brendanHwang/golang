/*
C언어와 다르게 break 없이도 다음 것으로 넘어가지 않게 설계됨

Switch 평가 순서
Switch case는 하나의 케이스라도 성공하면 멈추는 식으로 위에서부터 아래로 case를 평가합니다.

(예를 들어, 만약 i==0 라면 f를 호출하지 않습니다.)

switch i {
case 0:
case f():
}
주의: Go playground에서의 시간은 항상 2009-11-10 23:00:00 UTC에 시작하는 것으로 보여집니다. 이 값의 의미를 알아내는 것은 독자들의 활동으로 남겨져있습니다.

*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Tuesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 0:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

/*


조건이 없는 Switch
조건이 없는 Switch는 switch true 와 동일합니다.

이 구조는 긴 if-else 체인을 작성하는 데에 아주 깔끔한 방식일 수 있습니다.

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Day() == 31:
		fmt.Println("end of the month.")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

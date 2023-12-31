/*


Switch
switch 구문은 연속적인 if - else 구문을 사용하는 짧은 방안입니다. 그것은 값이 조건문과 같은 첫 번째 case를 실행합니다.

Go의 switch는 뒤이어 오는 모든 case를 실행하는 것이 아니라 오직 첫 번째로 선택된 케이스만을 실행한다는 점을 제외하고는 C나 C++, Java, Javascript, PHP와 유사합니다. 사실상 위의 언어들에서는 각각의 case의 마지막 부분에 break 구문이 필요하지만, Go에서는 자동으로 break가 제공됩니다. 다른 중요한 차이점은 Go의 switch case는 상수일 필요가 없으며 그 값들은 정수일 필요도 없습니다.

*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

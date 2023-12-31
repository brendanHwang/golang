// var
// var keyword는 함수 밖에서 사용할 수 있습니다 (함수 밖에서 선언을 해야하는 상황에서 사용)

package main

import (
	"fmt"
)

var a int = 10
var b, c, d bool

func main() {
	var i int
	fmt.Println(a, b, c, d, i)
}

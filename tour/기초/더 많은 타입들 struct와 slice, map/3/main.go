/*

For는 Go의 "while"
당신이 ; 을 생략할 수 있다는 점에서 C의 while 은 Go에서 for 로 쓰입니다.

*/

package main

import "fmt"

func main() {
	sum := 1
	for sum < 10000 {
		sum += sum
	}
	fmt.Println(sum)
}

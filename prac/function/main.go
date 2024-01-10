package main

import "fmt"

// 여러 인자 혹은 하나를 받아서 리턴
func greeting(names ...string) (gStr string) {

	gStr = "Hello,"
	for _, name := range names {
		gStr = fmt.Sprint(gStr, " ", name)
	}
	return gStr
}

func main() {
	fmt.Println(greeting("john", "doe", "jane"))
}

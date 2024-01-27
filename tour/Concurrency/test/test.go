package test

import (
	"log"
	"time"
)

func GoroutineTest() {
	time.Sleep(1 * time.Second)
	log.Println("Hello World!")

}

func GoroutineTest2(c chan string, i int) {
	c <- "Hello World! " + string(i)
}

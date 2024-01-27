package page2

import (
	"fmt"
	"time"
)

func Select() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "Hello from ch2"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		default:
			fmt.Println("No data received")
		}
		time.Sleep(time.Second * 1)
	}

}

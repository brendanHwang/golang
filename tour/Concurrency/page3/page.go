package page3

import "fmt"

func BuffedChannels() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 1
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

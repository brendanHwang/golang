package page2

import (
	"fmt"
	"sync"
)

// Channel 과 goroutine 을 이용한 blocking 그리고 WaitGroup 을 이용한 goroutine 의 종료를 기다리기
func WaitGroupChannelGoRoutine() {
	var wg sync.WaitGroup
	c := make(chan int, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		c <- 1
		c <- 2
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-c)
		fmt.Println(<-c)
	}()

	wg.Wait()

}

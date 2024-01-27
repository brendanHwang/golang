package page9

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string, wg *sync.WaitGroup) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
	wg.Done()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func SyncMutex() {

	wg := sync.WaitGroup{}
	c := SafeCounter{v: make(map[string]int)}

	wg.Add(1000)
	go func() {
		for i := 0; i < 1000; i++ {
			go c.Inc("somekey", &wg)
		}
	}()

	wg.Wait()
	fmt.Println(c.Value("somekey"))
}

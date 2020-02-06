package main

import (
	"fmt"
	"time"

	//	"sync"
	sync "github.com/aerth/syncdebug"
)

func main() {
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			mu.Lock()
			fmt.Println(x, "has the lock")
			<-time.After(time.Second * 2)
			mu.Unlock()
			fmt.Println(x, "gave up the lock")
			wg.Done()
		}(i)
	}
	wg.Wait()
}

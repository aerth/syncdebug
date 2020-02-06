package sync

import (
	"log"
	"runtime"
	"sync"
	"time"
)

type WaitGroup = sync.WaitGroup

type Mutex struct {
	mu sync.Mutex
}

var DefaultTimeout = time.Second

var Fatal = true

func (m *Mutex) Lock() {
	ch := make(chan (struct{}))
	_, caller, line, _ := runtime.Caller(1)
	go func(c chan (struct{})) {
		select {
		case <-time.After(DefaultTimeout):
			if Fatal {
				log.Fatalf("syncdebug: slow/no unlock at %s:%d", caller, line)
			}
			log.Printf("syncdebug: slow/no unlock at %s:%d", caller, line)
		case <-ch:
		}
	}(ch)
	m.mu.Lock()
	ch <- struct{}{}
}

func (m *Mutex) Unlock() {
	m.mu.Unlock()
}

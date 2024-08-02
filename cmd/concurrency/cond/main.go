package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	c     = sync.NewCond(&sync.Mutex{})
	ready = 5
	count = 0
)

func main() {
	for i := 0; i < 10; i++ {
		go Run(i)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()
	for {
		select {
		case <-time.After(time.Second):
			count++
			fmt.Println("count++: ", count)
			c.Broadcast()
		case <-ctx.Done():
			fmt.Println("7s timeout cancel")
			return
		}
	}
}

func Run(id int) {
	c.L.Lock()
	defer c.L.Unlock()
	for count != ready {
		c.Wait()
	}

	fmt.Printf("%d: Run\n", id)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	res := make(chan int, 10)

	go func() {
		defer close(jobs)
		for i := 0; i < 10; i++ {
			jobs <- i
		}
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, jobs, res, &wg)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for item := range res {
		fmt.Println(item)
	}
}

func worker(id int, jobs <-chan int, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("get job: %d\n", id)
		res <- job * job
	}
}

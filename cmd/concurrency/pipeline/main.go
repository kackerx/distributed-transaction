package main

import "fmt"

func gen() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	return ch
}

func add(in chan int) (out chan int) {
	out = make(chan int)
	go func() {
		defer close(out)
		for item := range in {
			out <- item + item
		}
	}()

	return out
}

func mul(in chan int) (out chan int) {
	out = make(chan int)
	go func() {
		defer close(out)
		for item := range in {
			out <- item * item
		}
	}()

	return out
}

func main() {
	for item := range mul(add(gen())) {
		fmt.Println(item)
	}
}

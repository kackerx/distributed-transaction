package main

import "fmt"

func main() {
	num := make(chan int, 1)
	letter := make(chan int, 1)

	go func() {
		num <- 0
	}()

	for {
		select {
		case i := <-num:
			fmt.Println(i)
			letter <- i
		case i := <-letter:
			fmt.Printf("%c\n", i+97)
			i++
			if i > 30 {
				return
			}
			num <- i
		}
	}
}

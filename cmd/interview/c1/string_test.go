package main

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	Name string
}

func TestStringNotAllEq(t *testing.T) {
	var v any = &User{"kk"}
	switch msg := v.(type) {
	case *User:
		fmt.Println(msg.Name)
	}
}

func TestName(t *testing.T) {
	ch := make(chan int, 100)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	close(ch)
	time.Sleep(time.Second * 5)
}

func TestHe(t *testing.T) {
	a := 1
	defer func() {
		fmt.Println("defer func:", a) // 2
	}()

	defer fmt.Println(a) // 1
	a = 2
	defer fmt.Println(a) // 2
}

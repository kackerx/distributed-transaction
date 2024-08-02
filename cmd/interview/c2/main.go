package main

import "fmt"

func foo(j int) (int, func() bool) {
	i := j + 1
	return i, func() bool {
		if i == 5 {
			return true
		}
		return false
	}

}
func main() {
	i, f := foo(3)

	fmt.Println(i, f())

	i++
	fmt.Println(i, f())
}

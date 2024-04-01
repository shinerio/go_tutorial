package main

import "fmt"

func closureFunc() func() {
	x := 100
	defer fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

func main() {
	f := closureFunc()
	f()
}

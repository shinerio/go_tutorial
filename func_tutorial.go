package main

import "fmt"

func main8() {
	a, b := test()
	fmt.Println(a, b)
}

func test() (a int32, b int32) {
	a, b = 1, 2
	return
}

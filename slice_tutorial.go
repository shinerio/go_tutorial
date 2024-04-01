package main

import "fmt"

func main5() {
	mySlice := make([]int, 3)
	newSlice := append(mySlice, 4)
	mySlice[0] = 1
	var newSlice2 = make([]int, 4)
	copy(mySlice, newSlice2)
	newSlice[0] = 2
	fmt.Println(mySlice)   // [0 0 0]
	fmt.Println(newSlice)  // [2 0 0 4]
	fmt.Println(newSlice2) // [0 0 0 0]
}

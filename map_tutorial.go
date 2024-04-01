package main

import (
	"fmt"
	"sync"
)

func main6() {
	var myMap sync.Map
	myMap.Store("a", 1)
	myMap.Store(2, "b")
	value, _ := myMap.Load("a")
	fmt.Println(value) // 1
	value, _ = myMap.Load(2)
	fmt.Println(value) // b
}

package main

import (
	"fmt"
	"time"
)

func main3() {
	// 无缓冲channel
	chan1 := make(chan int)
	chan2 := make(chan string)

	// 有缓冲channel
	chan3 := make(chan string, 5) // 缓冲区大小为5

	var data = "test"
	go func() {
		for {
			select {
			// <-chan1 是一个接收操作符号。如果chan1通道中有数据可以读取，那么就会执行这个case并丢弃读到的数据（仅当做一种触发信号），如果当前没有数据可读，则不会阻塞，而是继续检查下一个case。
			case <-chan1:
				fmt.Println("receive from chan1.")
			// 从chan3中读取数据赋值给tmp变量
			case tmp := <-chan3:
				fmt.Printf("receive from chan3, value %s\n", tmp)
			// chan2 <- data 是一个发送操作符号，其中data是要发送的数据。如果chan2通道可以立即发送数据,那么就会执行这个case下面的代码块。如果当前通道已满无法发送，则不会阻塞，而是继续检查下一个case。
			// 这里chan2未定义缓冲大小，如果没有其他goroutine从chan2中接收数据的会，这个发送动作会一直阻塞
			// 如果要在无接收者的情况下实现往chan2中发送数据，则最好给chan2定义缓冲大小
			case chan2 <- data:
				fmt.Println("send data to chan2")
			}
		}
	}()
	// 持续向chan1发送数据，以触发select中的case
	go func() {
		for i := 0; i < 3; i++ {
			chan1 <- i
		}
	}()

	// 持续从chan2中接收数据，以触发select中的case
	go func() {
		for i := 0; i < 3; i++ {
			var _ = <-chan2
		}
	}()

	// 持续向chan3发送数据，以触发select中的case
	go func() {
		for i := 0; i < 5; i++ {
			chan3 <- fmt.Sprintf("data%d", i)
		}
	}()

	time.Sleep(time.Second * 10)

	// 这里需要关闭chan，以防止goroutine泄漏
	close(chan1)
	close(chan2)
	close(chan3)
}

func main4() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 从一个已经关闭的chan中获取数据不会阻塞，会得到0，如下代码会打印 Received:1 Received: 2 Received: 3，之后一直死循环打印Received: 0
	//for {
	//	val := <-ch
	//	time.Sleep(time.Second)
	//	fmt.Println("Received:", val)
	//}

	// 从channel中获取数据时最好判断channel是否关闭了，除非你确定chan不会关闭
	for {
		// ok为bool类型，代表chan是否关闭
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Received:", val)
	}
}

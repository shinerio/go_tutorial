package main

import "fmt"

// 枚举
const (
	spring = 0
	summer = 1
	autumn = 2
	winter = 3
)

// constant block
const (
	one           = 1
	valueStillOne // value为1, 在常量块中,如果某个常量的值被省略,Go编译器会自动将这个常量的值设置为与上一个显式赋值的常量相同的值
	two           = 1
)

const (
	a = iota // 0
	b        // 1
	c        //2
	d = "ha" //独立值，iota += 1
	e        //“ha” iota += 1
	f = 100  //iota +=1
	g        //100  iota +=1
	h = iota //7,恢复计数
	i        //8
)

const (
	j = 1 << iota
	k = 3 << iota
	l // 3<<iota, iota为3，最后值12
	m // 3<<iota, iota为4，最后值24
)

func main1() {
	fmt.Println(i)
}

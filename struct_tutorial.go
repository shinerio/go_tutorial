package main

import "fmt"

// Shape 接口是一种抽象类型,它定义了一组方法签名,但不包含实现代码
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 实现了接口中声明的所有方法，那么它就被认为实现了该接口
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Square struct {
	edgeSize float64
}

func (r Square) Area() float64 {
	return r.edgeSize * r.edgeSize
}

func printInfo(shape Shape) {
	fmt.Println(shape.Area())      // 200
	fmt.Println(shape.Perimeter()) // 60
}

func main7() {
	myRectangle := Rectangle{
		Width:  10.0,
		Height: 20.0,
	}
	printInfo(myRectangle)
	//mySquare := Square{
	//	edgeSize: 10,
	//}

	// Cannot use 'mySquare' (type Square) as the type Shape Type does not implement 'Shape'
	//printInfo(mySquare)
}

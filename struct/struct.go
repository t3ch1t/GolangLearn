package main

import "fmt"

func main() {

	rect1 := rectangle{height: 20, width: 30}

	fmt.Println(rect1.width)
	fmt.Println(rect1.height)

	fmt.Println(rect1.area())

}

type rectangle struct {
	height float32
	width  float32
}

func (rect *rectangle) area() float32 {

	return rect.height * rect.width
}

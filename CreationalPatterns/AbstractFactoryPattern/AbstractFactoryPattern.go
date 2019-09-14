package AbstractFactoryPattern

import "fmt"

// 为形状创建一个接口
type Shape interface {
	draw()
}

// 创建实现接口的结构体
type Rectangle struct {
}

func (r Rectangle) draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}

type Square struct {
}

func (s Square) draw() {
	fmt.Println("Inside Square::draw() method.")
}

type Circle struct {
}

func (c Circle) draw() {
	fmt.Println("Inside Circle::draw() method.")
}

// 为颜色创建一个接口
type Color interface {
	fill()
}

type Red struct {
}

func (r Red) fill() {
	fmt.Println("Inside Red::fill() method.")
}

type Green struct {
}

func (r Green) fill() {
	fmt.Println("Inside Green::fill() method.")
}

type Blue struct {
}

func (r Blue) fill() {
	fmt.Println("Inside Blue::fill() method.")
}

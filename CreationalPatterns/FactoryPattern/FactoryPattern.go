package FactoryPattern

import (
	"fmt"
	"strings"
)

// 创建一个接口
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

// 创建一个工厂，生成基于给定信息的实例
type ShapeFactory struct {
}

func (s ShapeFactory) GetShape(shapeType string) Shape {
	if shapeType == "" {
		return nil
	} else if strings.ToLower(shapeType) == "circle" {
		return new(Circle)
	} else if strings.ToLower(shapeType) == "rectangle" {
		return new(Rectangle)
	} else if strings.ToLower(shapeType) == "square" {
		return new(Square)
	} else {
		return nil
	}
}

// 使用该工厂，通过传递类型信息来获取实体类的对象。
func Test() {
	shapeFactory := ShapeFactory{}

	shape1 := shapeFactory.GetShape("CIRCLE")
	shape1.draw()

	shape2 := shapeFactory.GetShape("RECTANGLE")
	shape2.draw()

	shape3 := shapeFactory.GetShape("SQUARE")
	shape3.draw()
}

/*
 通过一个结构体方法提供创建不同结构体实例，这些结构体都属于同一个类别（接口类型）
*/

package FactoryPattern

import "fmt"

// Factory and car interfaces
type CarFactory interface {
	makeCar() Car
}

type Car interface {
	getType() string
}

// Concrete implementations of the factory and car
type SedanFactory struct {
}

func (s SedanFactory) makeCar() Car {
	return Sedan{}
}

type Sedan struct {
}

func (s Sedan) getType() string {
	return "Sedan"
}

func Test() {
	factory := SedanFactory{}
	car := factory.makeCar()
	fmt.Println(car.getType())
}

/*
refer: https://en.wikipedia.org/wiki/Factory_method_pattern
使用工厂模式而不是直接实例化结构体可以更好地遵循SOLID（https://en.wikipedia.org/wiki/SOLID）原则
*/

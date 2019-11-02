package Visitor

import (
	"fmt"
)

type Visitor interface {
	visitCar(element Element)
	visitTree(element Element)
}

type Element interface {
	accept(visitor Visitor)
	Name() string
}

type Car struct {
	name string
}

func (c Car) Name() string {
	return c.name
}

func (c Car) accept(visitor Visitor) {
	visitor.visitCar(c)
}

type Tree struct {
	name string
}

func (t Tree) Name() string {
	return t.name
}

func (t Tree) accept(visitor Visitor) {
	visitor.visitTree(t)
}

type Scan struct {
	speed string
}

func (s Scan) visitCar(element Element) {
	fmt.Printf("Scan %s with speed %s\n", element.Name(), s.speed)
}

func (s Scan) visitTree(element Element) {
	fmt.Printf("Scan %s with speed %s\n", element.Name(), s.speed)
}

type Overlook struct {
}

func (o Overlook) visitCar(element Element) {
	fmt.Printf("Overlook %s\n", element.Name())
}

func (o Overlook) visitTree(element Element) {
	fmt.Printf("Overlook %s\n", element.Name())
}

func Test() {
	elements := []Element{Car{"theCar"}, Tree{"theTree"}}
	visitors := []Visitor{Scan{"5m/s"}, Overlook{}}

	for _, visitor := range visitors {
		for _, e := range elements {
			e.accept(visitor)
		}
	}
}

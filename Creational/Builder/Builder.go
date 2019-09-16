package Builder

import (
	"fmt"
	"log"
)

// product
type Fruit struct {
	name, weight, desc string
}

// builder interface
type FruitBuilder interface {
	SetWeight(w string)
	MakeFruit() (*Fruit, error)
}

// apple builder
type AppleBuilder struct {
	myWeight string
}

func (a *AppleBuilder) SetWeight(w string) {
	a.myWeight = w
}

func (a AppleBuilder) MakeFruit() (*Fruit, error) {
	return &Fruit{"apple", a.myWeight, fmt.Sprintf("I'm an %s, %s weighted", "apple", a.myWeight)}, nil
}

// banana builder
type BananaBuilder struct {
	myWeight string
}

func (b *BananaBuilder) SetWeight(w string) {
	b.myWeight = w
}

func (b BananaBuilder) MakeFruit() (*Fruit, error) {
	return &Fruit{"banana", b.myWeight, fmt.Sprintf("I'm a %s, %s weighted", "banana", b.myWeight)}, nil
}

// director
type machine struct {
}

func (m machine) ProduceFruit(builder FruitBuilder) *Fruit {
	builder.SetWeight("3kg")
	f, err := builder.MakeFruit()
	if err != nil {
		log.Fatal("ProduceFruit:", err)
	}
	return f
}

func Test() {
	m := machine{}
	apple := m.ProduceFruit(&AppleBuilder{})
	fmt.Println(apple.desc)
	banana := m.ProduceFruit(&BananaBuilder{})
	fmt.Println(banana.desc)
}

/*
refer: http://blog.ralch.com/tutorial/design-patterns/golang-builder/
1. 通过定义不同的builder控制实例化的过程
2. 通过接口变量调用具体方法，类似虚函数
*/

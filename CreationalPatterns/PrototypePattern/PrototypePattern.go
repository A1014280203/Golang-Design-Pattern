package PrototypePattern

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type concretePrototype struct {
	Name string
}

func (c *concretePrototype) SetName(name string) {
	c.Name = name
}

func (c concretePrototype) Clone() concretePrototype {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(c)
	if err != nil {
		log.Fatal("encode:", err)
	}
	dec := gob.NewDecoder(&buf)
	var _c concretePrototype
	err = dec.Decode(&_c)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return _c
}

func Test() {
	var c1 concretePrototype
	c1.SetName("c1")
	c2 := c1.Clone()
	fmt.Println(c1.Name == c2.Name)
	fmt.Printf("%p\n", &c1)
	fmt.Printf("%p\n", &c2)
}

/*
refer: http://blog.ralch.com/tutorial/design-patterns/golang-prototype/
refer: https://golang.org/pkg/encoding/gob/
所谓原型模式，就是深拷贝已经存在的实例
Golang没有内建方法支持结构体的深拷贝，但可以使用内建序列化方法gob实现
*/

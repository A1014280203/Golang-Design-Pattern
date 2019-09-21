package Facade

import "fmt"

type TyreMaker struct {
	total int
}

func (t *TyreMaker) Make() {
	fmt.Println("making", t.total, "tyres")
}

type WindowMaker struct {
	total int
}

func (w *WindowMaker) Make() {
	fmt.Println("making", w.total, "windows")
}

type SkeletonMaker struct {
	total int
}

func (s *SkeletonMaker) Maker() {
	fmt.Println("making", s.total, "skeleton")
}

type CarMaker struct {
	brand string
}

func (c *CarMaker) Make() {
	tyreMaker := TyreMaker{4}
	windowMaker := WindowMaker{4}
	skeletonMaker := SkeletonMaker{1}
	fmt.Println(c.brand, "car is making")
	tyreMaker.Make()
	windowMaker.Make()
	skeletonMaker.Maker()
	fmt.Println(c.brand, "car made")
}

func Test() {
	carMaker := CarMaker{"bester"}
	carMaker.Make()
}

/*
refer: https://www.runoob.com/design-pattern/design-pattern-tutorial.html
通过一个接口操作整个处理过程
*/

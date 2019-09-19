package Decorator

import (
	"fmt"
)

type Ultraman struct {
	restEnergy int
}

func (u Ultraman) Ray(energy int) {
	var i = 0
	for ; i < energy-1 && i < u.restEnergy-1; i++ {
		fmt.Printf("=")
	}
	u.restEnergy -= energy
	if i > 0 {
		fmt.Println(">")
	} else {
		fmt.Println("No power now")
	}
}

// 函数签名和非结构体方法一样
func DescribeRay(f func(e int)) func(e int) {
	return func(e int) {
		fmt.Println("At last", e, "power")
		f(e)
		fmt.Println("Booooom!")
	}
}

func Test() {
	var u Ultraman
	u.restEnergy = 0
	n := DescribeRay(u.Ray)
	n(30)
}

/*
refer: https://github.com/bvwells/go-patterns/blob/master/structural/decorator.go
工具：函数残参数+函数闭包
*/

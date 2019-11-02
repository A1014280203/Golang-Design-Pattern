package Strategy

import "fmt"

type Do func(num1, num2 int) int

type Calculator struct {
	do Do
}

func (c Calculator) calculate(num1, num2 int, op string) (r int) {
	switch op {
	case "+":
		r = Do(sum)(num1, num2)
	case "-":
		r = Do(sub)(num1, num2)
	}
	return r
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func sub(num1, num2 int) int {
	return num1 - num2
}

func Test() {
	num1 := 1
	num2 := 2
	c := Calculator{}
	fmt.Println(c.calculate(num1, num2, "+"))
	fmt.Println(c.calculate(num1, num2, "-"))
}

/*
和State模式的相同点，都是将行为的改变封装在里面。
不同之处：
1. 外部使用者对内部调用的方法是用感知或者期望的，比如这里的加法，调用知道且期望调用加法方法
2. 内部调用何种方法不是状态控制的，是由包含状态的对象控制的。在这里是Calculator
*/

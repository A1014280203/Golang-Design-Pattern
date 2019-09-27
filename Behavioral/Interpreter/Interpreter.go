package Interpreter

import (
	"fmt"
	"strings"
)

type Context string

type Expression interface {
	Interpret() int
}

type Compare struct {
	a Context
	b Context
}

func (c Compare) Interpret() int {
	return strings.Compare(string(c.a), string(c.b))
}

type Contain struct {
	a Context
	b Context
}

func (c Contain) Interpret() int {
	if strings.Contains(string(c.a), string(c.b)) {
		return 1
	} else {
		return 0
	}
}

func CreateExpression(choice, a, b Context) Expression {
	switch choice {
	case "compare":
		return Compare{a, b}
	case "contain":
		return Contain{a, b}
	default:
		return nil
	}
}

func Test() {
	var a Context = "banana"
	var b Context = "nana"
	exp := CreateExpression("compare", a, b)
	fmt.Println(exp.Interpret())
	exp = CreateExpression("contain", a, b)
	fmt.Println(exp.Interpret())
}

/*
refer:
1. https://en.wikipedia.org/wiki/Interpreter_pattern
2. https://github.com/sevenelevenlee/go-patterns/blob/master/26-interpreter-pattern/interpreter.go
属于结构模式中的Composite模式，
*/

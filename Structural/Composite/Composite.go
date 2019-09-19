package Composite

import "fmt"

type Member interface {
	Report()
}

type Student int

func (s Student) Report() {
	fmt.Println("I'm student", s)
}

// School
type School struct {
	name string
}

// student
type Students struct {
	s []Student
}

// camp group
type Group struct {
	School
	Students
}

func (g *Group) Report() {
	fmt.Println("from", g.name, "school")
	for _, v := range g.Students.s {
		v.Report()
	}
}

func NewGroup() *Group {
	s := make([]Student, 5)
	var i Student
	for i = 0; i < 5; i++ {
		s[i] = i
	}
	return &Group{School{"W"}, Students{s}}
}

func Test() {
	g := NewGroup()
	g.Report()
}

/*
refer: https://github.com/sevenelevenlee/go-patterns/blob/master/13-composite-pattern/composite.go
将Student的Report方法包装在Group的Report方法之下
*/

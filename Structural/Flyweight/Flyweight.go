package Flyweight

import "fmt"

type SchoolInfo string

type Student struct {
	id         int
	schoolInfo *SchoolInfo
}

func (s *Student) SetSchool(schoolInfo *SchoolInfo) {
	s.schoolInfo = schoolInfo
}

func (s *Student) PrintInfo() {
	fmt.Printf("I'm %d, from %s\n", s.id, *s.schoolInfo)
}

func Test() {
	var info SchoolInfo
	// assume this cost a a lot, so that only one kept
	info = "yalasuo"
	//
	students := make([]Student, 10)
	for i := range students {
		students[i] = Student{i, &info}
	}
	for _, v := range students {
		v.PrintInfo()
	}
}

/*
refer: https://en.wikipedia.org/wiki/Flyweight_pattern
很多实例都需要同一个数据，则将其拆出，并通过实例的set方法将其赋值给实例
*/

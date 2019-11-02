package Observer

import "fmt"

// Subject
type Teacher struct {
	students      []*Student
	todayHomework string
}

func (t *Teacher) appendStudent(kid *Student) {
	t.students = append(t.students, kid)
}

func (t *Teacher) assignmentHomework(task string) {
	t.todayHomework = task
	for _, student := range t.students {
		// notify homework
		student.getHomework(t.todayHomework)
	}
}

// Observer
type Student struct {
	name string
}

func (s Student) getHomework(homework string) {
	fmt.Printf("%s's homework is %s\n", s.name, homework)
}

func Test() {
	teacher := Teacher{}
	John := Student{"John"}
	Mike := Student{"Mike"}
	teacher.appendStudent(&John)
	teacher.appendStudent(&Mike)
	teacher.assignmentHomework("reading")
}

/*
observe模式，包含observer和subject两部分。subject状态发生变化时，会通知依赖其的observer。
wiki里面说这种模式适于事件处理，应该是指回调
*/

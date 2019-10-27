package Iterator

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type School struct {
	names []string
	index int
}

func (s *School) HasNext() bool {
	if s.index < len(s.names) {
		return true
	}
	return false
}

func (s *School) nextIndex() {
	s.index++
}

func (s *School) Next() interface{} {
	if s.HasNext() {
		defer s.nextIndex()
		return s.names[s.index]
	}
	panic("")
}

func Test() {
	schools := School{names: []string{"A", "B"}}
	var names Iterator = &schools
	for names.HasNext() {
		fmt.Println(names.Next())
	}
}

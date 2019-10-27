package Iterator

import (
	"errors"
	"fmt"
)

type Iterable interface {
	Next() (interface{}, error)
}

type Teachers struct {
	names []string
	index int
}

func (t *Teachers) Next() (name interface{}, err error) {
	if t.index < len(t.names) {
		name = t.names[t.index]
		t.index++
		err = nil
	} else {
		name = ""
		err = errors.New("StopIteration")
	}
	return name, err
}

func Test2() {
	teachers := Teachers{names: []string{"A", "B"}}
	var names Iterable = &teachers
	for name, err := names.Next(); err == nil; name, err = names.Next() {
		fmt.Println(name)
	}
}

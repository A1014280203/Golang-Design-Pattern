package Iterator

import "fmt"

type Iter interface {
	GetArray() []Pair
}

type Students struct {
	names []string
	id    []int
}

type Pair struct {
	string
	int
}

func (s Students) GetArray() []Pair {
	var r = make([]Pair, len(s.names))

	for i := 0; i < len(s.names); i++ {
		r[i] = Pair{s.names[i], s.id[i]}
	}
	return r
}

func Test3() {
	students := Students{[]string{"A", "B"}, []int{1, 2}}
	var iter Iter = students
	for _, v := range iter.GetArray() {
		fmt.Println(v)
	}
}

/*
Python风格的迭代器，即返回一个可迭代实例，在这里是用的Go的数组
*/

package Filter

import "fmt"

type Person struct {
	id   int
	name string
}

func GroupById(persons []Person) (odd []Person, even []Person) {
	for _, v := range persons {
		if (v.id % 2) == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return odd, even
}

func Test() {
	persons := make([]Person, 4)
	persons[0] = Person{1, "askdoiaui"}
	persons[1] = Person{2, "vwoejhoi"}
	persons[2] = Person{3, "cbqyet"}
	persons[3] = Person{4, "nucetwqy"}
	oddId, evenId := GroupById(persons)
	for _, v := range oddId {
		fmt.Printf("%d %s	", v.id, v.name)
	}
	fmt.Println()
	for _, v := range evenId {
		fmt.Printf("%d %s	", v.id, v.name)
	}
}

/*
refer: https://www.runoob.com/design-pattern/filter-pattern.html
*/

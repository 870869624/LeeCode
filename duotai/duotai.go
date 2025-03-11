package main

import (
	"fmt"
)

type one interface {
	make() error
}

type all interface {
	one
}

func do(a all) error {
	return a.make()
}

type student struct {
	eat  string
	sing string
}

type person struct {
	student
	rus string
}

func (s *student) make() error {
	fmt.Println(s)
	return nil
}
func (p *person) make() error {
	fmt.Println(p)
	return nil
}

func main() {
	s1 := &student{
		eat:  "好吃的",
		sing: "紫金花",
	}
	do(s1)
	p1 := &person{
		student{
			eat:  "ssss",
			sing: "ssss",
		},
		"ssssss",
	}
	do(p1)

}

package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, i am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La la la...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, i am %s, i work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func blankinterface() {
	type blankinter interface {

	}

	var i blankinter

	a := 1

	i = a
	fmt.Println(i)

}

func CreateStudent() Men {
	return Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	fmt.Println(mike)
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	//sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}


	var i Men
	i = mike
	i.SayHi()
	i.Sing("November rain")

	i = Tom
	i.SayHi()
	i.Sing("Born to be wild")

	x := make([]Men, 3)
	x[0], x[1], x[2] = mike, paul, Tom

	for _, value := range x {
		value.SayHi()
	}

	blankinterface()

	a := CreateStudent()
	a.SayHi()

}




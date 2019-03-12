package main

import "fmt"

type Pet interface {
	Name() string
	Age() uint8
}

type Dog struct {
	name string
	age uint8
}

func (mydog Dog) Name() string {
	return mydog.name
}

func (mydog Dog) Age() uint8 {
	return mydog.age
}

func main() {
	myDog := Dog{"Little D", 3}
	_, ok1 := interface{}(&myDog).(Pet)
	_, ok2 := interface{}(myDog).(Pet)
	fmt.Printf("%v, %v\n", ok1, ok2)
}

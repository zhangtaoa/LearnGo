package main

import "fmt"

type Animal interface {
	Grow()
	Move(string) string
}

type Cat struct {
	Name string
	Age uint8
	Addr string 
}

func (cat *Cat) Grow() {
	cat.Age++
}

func (cat *Cat) Move(newaddr string) string {
	var oldaddr = cat.Addr
	cat.Addr = newaddr
	return oldaddr
}

func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	animal, ok := interface{}(&myCat).(Animal)
	fmt.Printf("%v, %v\n", ok, animal)
}

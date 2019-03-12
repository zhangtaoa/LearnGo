package main

import "fmt"

type Person struct {
	Name string
	Gender string
	Age uint8
	Addr string
}

func (person *Person) Move(newaddr string) string{
	var oldaddr = person.Addr
	person.Addr = newaddr
	return oldaddr
}

func main() {
	p := Person{"Robert", "Male", 33, "Beijing"}
	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s. \n", p.Name, oldAddress, p.Addr)

}

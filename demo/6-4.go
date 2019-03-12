package main

import "fmt"

type MyInt struct {
	n int
}

func (myint *MyInt) incr() {
	myint.n++
}

func (myint MyInt) decr() {
	myint.n--
}

func main() {
	myint := MyInt{}
	myint.incr()
	myint.incr()
	myint.decr()
	myint.decr()
	myint.incr()
	fmt.Println(myint.n == 1)
	fmt.Println(MyInt{})

}

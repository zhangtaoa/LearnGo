package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	for {
		n, ok  := <-c
		if !ok {
			break
		}
		fmt.Println(n)
	}
}

func creatework() chan int{
	c := make(chan int)
	go worker(c)
	return c
}
func chanDemo() {
	//c := make(chan int)
	//go worker(c)
	c := creatework()
	c <- 1
	c <- 2
	c <- 3
}

func chanClose() {
	c := make(chan int)
	go worker(c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
}



func main(){
	//chanDemo()
	chanClose()
	time.Sleep(time.Millisecond)
}
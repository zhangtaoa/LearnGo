package main

import (
	"fmt"
)

func test() {
	ch1 := make(chan int, 1)	
	ch2 := make(chan int, 1)	
	ch1 <- 1111
	ch2 <- 2222
	select {
		case e1 := <-ch1:
			fmt.Printf("1th e1 = %v.\n",e1)
		case e2 := <-ch2:
			fmt.Printf("2th e2 = %v.\n",e2)
		default:
			fmt.Println("no data")
	}
}

func test1() {
	ch4 := make(chan int, 1)
	for i := 0; i < 4; i++ {
		select {
			case e, ok := <-ch4:
				if !ok {
					fmt.Println("End.")
					return
				}
				fmt.Println(e)
				close(ch4)
			default:
				fmt.Println("No Data!")
				ch4 <- 1
		}
	}
}

func main() {
	test1()
}

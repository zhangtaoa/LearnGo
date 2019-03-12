package main

import (
	"fmt"
	"time"
)

func main() {
	for i :=0 ; i < 100 ; i++ {
		go func(i int){
			fmt.Printf("my name is %d\n", i)
		}(i)
	}

	time.Sleep(time.Minute)
}


package main

import "fmt"

func main() {
	var number int = 5
	if number += 4; 10 > number {
		number += 27
		number += 3
		fmt.Println(number)		
	} else if 10 < number {
		number -= 2
		fmt.Println(number)
	}
	fmt.Println(number)

}

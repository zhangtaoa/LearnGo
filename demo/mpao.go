package main

import "fmt"

func main(){
	my_list := [5]int{1 ,2 ,3 ,4 ,5}
	fmt.Println(my_list)
	fmt.Printf("my_list len is %d\n", len(my_list))

	for _, value := range my_list {
		fmt.Println(value)
	}

	x1 := make([]int, 5)
	fmt.Println(x1)

	fmt.Println(x1[1:3], len(x1[1:3]), cap(x1[1:3]))

	var x2 [5]int

	for i, _ := range x2{
		x2[i] = i*100
	}

	fmt.Println(x2)
}

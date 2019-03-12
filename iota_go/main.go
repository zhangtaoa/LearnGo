package main

import "fmt"

// 自增变量

const (
	age  int = 1 * iota
	age1
	age3
)

func main()  {
	fmt.Println(age)
	fmt.Println(age1)
	fmt.Println(age3)
}

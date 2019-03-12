package main

import (
	"fmt"
	"math/rand"
)

func test1() {
	name := "Golang"
	switch name {
		case "Golang":
			fmt.Println("Golang")
		case "Rust":
			fmt.Println("Rust")
		default:
			fmt.Println("default")
	}
}

func test2() {
	names := []string{"Golang", "Java", "Rust", "C"}
	switch name := names[0]; name {
		case "Golang":
			fmt.Println("Golang")
		case "Java":
			fmt.Println("Java")
		case "Rust":
			fmt.Println("Rust")
		case "C":
			fmt.Println("C")
	}

}

func test3() {
	v := 11
	switch i := interface{}(v).(type) {
		case int, int8, int16, int32, int64:
			fmt.Printf("%d,%T",i, i)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("%d,%T",i, i)
	}
}

func test4() {
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	fmt.Println(ia[rand.Intn(4)/4 + 3])
	switch v := ia[rand.Intn(4):4]; interface{}(v).(type){
		case []interface{}:
			fmt.Println("CASE A")
		case byte:
			fmt.Println("CASE B")
		default:
			fmt.Println("default")
	}
}

func main() {
	test4()
}

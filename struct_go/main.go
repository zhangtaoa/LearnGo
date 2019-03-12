package main

import (
	"fmt"
	"reflect"
)

type Persion struct {
	name string
	age int
	addr string
}

func main() {
	staff_zt := Persion{"zhangtao", 30, "bj"}
	fmt.Println(staff_zt)
	fmt.Printf("name: %s\nage: %d\naddr: %s\n\n", staff_zt.name, staff_zt.age, staff_zt.addr)

	fmt.Printf("%v",staff_zt)

	v := reflect.ValueOf(staff_zt)
	fmt.Println(v)
	fmt.Println(v.Type())

}

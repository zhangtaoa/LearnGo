package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return ioutil.ReadAll(file)
}

func test1() {
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(2)
	}()

	defer func() {
		fmt.Println(3)
	}()
}

func test2() {
	for i := 1; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func test3() {
	f := func (i int) int {
		fmt.Printf("%d\n", i)
		return i*10
	}

	for i :=1; i < 5; i++ {
		defer fmt.Printf("%d\n", f(i))
	}
}

func test4() {
	for i := 1; i < 5; i++ {
		defer func(){ 
			fmt.Println(i)
		}()
	}
}

func test5() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d ", fab(i));defer fmt.Printf("%d ", fab(i))
	}
}	

func fab(i int) int {
		if i == 0 {
			return 0
		}
		if i == 1 {
			return 1
		}
		return fab(i-1) + fab(i-2)
	}

func main() {
	//fmt.Println(readFile("/Users/tao.zhang/mygo/1.go"))
	//test1()
	//test2()
	test5()
}

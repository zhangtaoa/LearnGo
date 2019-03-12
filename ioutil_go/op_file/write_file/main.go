package main

import (
	"bufio"
	"fmt"
	"os"
)

const file_name = "./a.txt"

func create_file() (*os.File) {
	file, err :=os.Create(file_name)
	if err != nil{
		fmt.Println("Create err:", err)
	}

	return file
}

func main() {
	file := create_file()
	rw := bufio.NewWriter(file)
	rw.WriteString("aaaaaaaaaaaaaaaaaaaa")
	rw.Flush()
}

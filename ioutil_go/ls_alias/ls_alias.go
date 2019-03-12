package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	dir_path := "."
	files , err := ioutil.ReadDir(dir_path)
	if err != nil {
		fmt.Println("ReadDir error: ", err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}



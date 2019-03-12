package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	url := "http://127.0.0.1:8080/hello"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get err:",err)
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll err:", err)
	}

	fmt.Println(string(content))
}

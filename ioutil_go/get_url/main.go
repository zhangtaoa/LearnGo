package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://www.baidu.com"

	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get err ", err)
	}

	defer rsp.Body.Close()

	content, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(content))
}

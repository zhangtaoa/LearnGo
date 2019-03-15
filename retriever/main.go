package main

import (
	"LearnGo/retriever/mock"
	"fmt"
)

type retriever interface {
	Get(url string) string
}

func download(r retriever) string {
	return r.Get("www.baidu.com")
}

func main() {
	var r retriever
	r = mock.Retriever{"zhangtao"}
	fmt.Println(download(r))
}



package main

import (
	"LearnGo/retriever/mock"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"This is zhangtao test"}
	fmt.Printf("%T %v\n", r, r)
	fmt.Println(download(r))

	r1 := mock.Retriever{"This is zhangtao test"}
	fmt.Println(r1.Content)

	//r = real.Retriever{}
	//fmt.Println(download(r))
}

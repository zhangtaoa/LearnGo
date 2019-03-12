package mock

import "fmt"

type Retriever struct {
	Content string
}

func (r Retriever) Get(url string) string {
	fmt.Println(url)
	return r.Content
}

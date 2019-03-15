package mock

type Retriever struct {
	Content string
}

func (r Retriever) Get(url string) string {
	return "This is a mock!"
}
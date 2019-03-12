package main

import (
	"fmt"
	"net/http"
	"io"
)

func HelloServer(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.Header)
	io.WriteString(w, "hello world")
}

func main() {
	http.HandleFunc("/hello", HelloServer)

	http.ListenAndServe(":8080", nil)
}

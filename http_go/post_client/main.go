package post_client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func mian() {
	send_msg := "This is my first client."
	rsp, err := http.Post("http://127.0.0.1:8080/hello",
		"application/text;charset=utf-8", bytes.NewBufferString(send_msg))
	if err != nil {
		print("http post err:", err)
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(body))

}
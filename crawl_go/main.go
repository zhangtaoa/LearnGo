package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	rsp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		panic("rsp not 200")
	}

	msg, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s", msg)
	printcity(string(msg))

}

func printcity(context string){
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)" data-v-5e16505f>([^<]+)</a>`)

	msgs := reg.FindAllStringSubmatch(context, -1)

	for _, msg := range msgs{
		fmt.Printf("city:%s, url:%s\n", msg[2], msg[1])
	}

	fmt.Println(len(msgs))
}

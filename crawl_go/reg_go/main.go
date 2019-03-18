package main

import (
	"fmt"
	"regexp"
)



func main() {
	const msg = `this is my email tao.zhang@moji.com.cn 
		ding aaa@bbb.com 
		dd   ccc@iii.org`
	reg, err := regexp.Compile(`([.a-zA-Z0-9]+)@([a-zA-Z0-9]+)([\.a-zA-Z0-9]+)`)
	if err != nil {
		panic(err)
	}

	res := reg.FindAllStringSubmatch(msg, -1)

	fmt.Println(res)

}

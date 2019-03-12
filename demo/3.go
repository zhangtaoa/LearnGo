package main

import (
	"fmt"
	"strings"
)

func main() {

	var mystr string = "abcdabc"
	var max_sub_len int
	mystr_len := len(mystr)
	var sub_str string
	i := 0
	var j int
	for i=0;i<mystr_len-1;i++ {
		for j=i+1;j<mystr_len;j++ {
		   	sub_str = mystr[i:j]
		   	sub_str = mystr[i:j]
			if strings.Contains(sub_str, mystr[j:j+1]) {
					max_sub_len = len(sub_str)
					fmt.Println(max_sub_len)
					break;
				}
		}
		
	}

}

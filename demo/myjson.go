package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	my_dict := [2]map[string]string{{"name": "zhangtao", "age": "18"},
		{"name": "zhangtao", "age": "38"},}
	fmt.Println(my_dict)

	for key, value := range my_dict {
		fmt.Println(key, value)
	}

	value, _ :=json.Marshal(my_dict)
	fmt.Println(value)
}

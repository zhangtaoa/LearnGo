package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type PersonAll struct {
	Personall []PersonInfo `json:"person"`
}


type PersonInfo struct {
	Name   string    `json:"aaa"`
	Age    int        `json:"age"`
}

func decode_test() {
	JsonMsg := `{"person":[{
	"aaa": "test1",
	"age": 17
},
{
	"aaa": "test2",
	"age": 18
}
]}`

	var info PersonAll

	err := json.Unmarshal([]byte(JsonMsg), &info)
	if err != nil {
		panic(err)
	}
	fmt.Println(info)


	var info1 PersonAll
	err1 := json.NewDecoder(strings.NewReader(JsonMsg)).Decode(&info1)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(info1)

}


func encode_test() {
	var personall PersonAll

	var personinfo []PersonInfo
	for i:=0; i<3; i++ {
		personinfo_temp := PersonInfo{fmt.Sprintf("test%d", i), i}
		personinfo = append(personinfo, personinfo_temp)
	}

	personall = PersonAll{personinfo}
	fmt.Println(personall)

	msg, err := json.Marshal(personall)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", msg)


	var s bytes.Buffer
	err1 := json.NewEncoder(&s).Encode(&personall)
	if err1 != nil {
		panic(err1)
	}

	fmt.Printf("msg:%s\n",s.String())


}


func main()  {
	//decode_test()
	encode_test()

}

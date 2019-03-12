package main

import "fmt"

type IFace interface {
	SetSomeField(newValue string)
	GetSomeField() string
}

type Implementation struct {
	someField string
}

func (i *Implementation) SetSomeField(value string){
	i.someField = value
}

func (i *Implementation) GetSomeField() string {
	return i.someField
}

func Create() *Implementation {
	return &Implementation{"hello"}
}

func main()  {
	var i IFace
	i = Create()
	fmt.Println(i.GetSomeField())
	i.SetSomeField("shit")
	fmt.Println(i.GetSomeField())
}

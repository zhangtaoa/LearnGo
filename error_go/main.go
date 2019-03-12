package main

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (err MyError) Error() string{
	return fmt.Sprintf("%v : %v", err.When, err.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"this my error",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}


	nodataerror := errors.New("No Data!")
	fmt.Println(nodataerror.Error())

	//myerr := MyError{
	//	time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
	//	"this my error",
	//}
	//
	//fmt.Println(myerr.Error1())
}
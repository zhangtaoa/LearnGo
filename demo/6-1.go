package main

import (
	"fmt"
		"strconv"
		"sync/atomic"
)

type EmployeeIdGenerator func(company string, department string, sn uint32) string

var company = "Gophers"
var sn uint32

func generateId(generator EmployeeIdGenerator, department string) (string, bool){
	if generator == nil {
		return "", false
	}

	newSn := atomic.AddUint32(&sn, 1)
	return generator(company, department, newSn), true
}

func appendSn(firstPart string, sn uint32) string {
	return firstPart + strconv.FormatUint(uint64(sn), 10)
}

func main() {
	var generator EmployeeIdGenerator
	generator = func(company string, department string, sn uint32) string{
			return appendSn(company+"-"+department+"-", sn)
		}
	fmt.Println(generateId(generator, "RD"))
}

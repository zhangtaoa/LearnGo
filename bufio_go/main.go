package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadBuf() {
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Print("Open file err:", err)
	}
	defer file.Close()

	rb := bufio.NewReader(file)
	fmt.Println(rb.Discard(1))
	fmt.Println(rb.ReadString('\n'))
}

func WriteBuf(){
	file, err := os.Create("./bufio_go/test.txt")
	if err != nil {
		fmt.Println("Create file err:", err)
		return
	}

	w_buf := "This is firsr line"

	//count, err := file.WriteString(w_buf)
	//if err != nil {
	//	fmt.Println("WriteString err:", err)
	//	return
	//}
	//fmt.Println("WriteString succ count:%d", count)
	wb := bufio.NewWriter(file)
	wb.WriteString(w_buf)
	wb.Flush()

}

func main() {
	//ReadBuf()
	WriteBuf()
}

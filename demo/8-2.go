package main

import (
	"io"
	"os"
	"fmt"
	"errors"
	"bytes"
	"bufio"
)

func test1(path string) (error) {

	if path == "" {
		return errors.New("The Parameter is invalid!")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	br := bufio.NewReader(file)
	var buf bytes.Buffer
	for {
		ba, isPrefix, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		buf.Write(ba)
		if !isPrefix {
			buf.WriteByte('\n')
		}
	}
	fmt.Printf("%s",buf)
	
	return nil
}

func main() {
	//fmt.Printf("%s", test1(""))
	fmt.Println(test1("/Users/tao.zhang/mygo/1.go"))

}

package read_file

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// get path
	fmt.Println(os.Getwd())

	// read file 01
	file_name := "./ioutil_go/op_file/main.go"
	//content, err := ioutil.ReadFile(file_name)
	//if err != nil {
	//	fmt.Println("read file err,", err)
	//}
	//fmt.Println(string(content))

	// read file 02
	//file, err := os.Open(file_name)
	//if err != err {
	//	fmt.Println("Open err :", err)
	//}
	//
	//for {
	//	data := make([]byte, 100)
	//	count, err := file.Read(data)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	//}

	// read file 03
	if file, err := os.Open(file_name); err == nil {
		defer file.Close()
		//reader := bufio.NewReader(file)

		// Read
		//for {
		//	data := make([]byte, 100)
		//	if count, err := reader.Read(data);err == nil {
		//		fmt.Println(string(data[:count]))
		//	} else {
		//		break
		//	}
		//
		//}

		// ReadString
		//for {
		//	str, err := reader.ReadString('\n')
		//	if err != nil {
		//		fmt.Println("ReadString err:", err)
		//		break
		//	}
		//	fmt.Print(str)
		//}

		// Scanner
		scan := bufio.NewScanner(file)
		fmt.Println(scan.Text())

	}
}

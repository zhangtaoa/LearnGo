package main

import (
	"io"
	"log"
	"os"
)

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	log.Printf("zhangtao")

	file, err := os.OpenFile("./log_go/main.log", os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	m_wido := io.MultiWriter(file, os.Stdout)

	info_log := log.New(m_wido, "INFO:", log.Ldate | log.Ltime | log.Llongfile)
	info_log.Printf("op err, err_info:%s", "not found file , net")

}

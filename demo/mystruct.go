package demo

import "fmt"

type AllInfo interface {
	getName() string
}

type Info struct {
	name string
	age int
	addr string
}

func (info *Info) getName() string {
	return info.name
}

func getallName(allinfo AllInfo) string {
	fmt.Println(allinfo.getName)
}


func main() {
	zt_info := Info{"zhangtao", 30,"bj"}

	//fmt.Println(zt_info)
	//fmt.Println(zt_info.getName())

	var all_info AllInfo
	all_info = &zt_info
	fmt.Println(all_info.getName())

}
package main
import "fmt"

func init(){
	fmt.Println("init...")
}

func init(){
	fmt.Println("init22222")
}

func test(){
	a := 3
 	fmt.Println(a)
}

func main(){
	fmt.Println("main...")
        test()
}

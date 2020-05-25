package main
import "fmt"

func main() {
	//golang变量使用方式1
	var i int
	fmt.Println("i=", i)
	// i= 0

	//golang变量使用方式2
	var num = 10.11
	fmt.Println("num=", num)

	//golang变量使用方式3
	name := "Tom"
	// := 的 : 不能省略
	// 等价于 var name string 加上 name = "tom"
	fmt.Println("name=", name)
}

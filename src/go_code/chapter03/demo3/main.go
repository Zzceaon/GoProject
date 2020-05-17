package main
import "fmt"

// 全局变量
// var n1 = 100
// var n2 = 200
// var name = "jack"
// 一次性声明
var (
	n8 = 800
	n9 = 900
	name1 = "mary"
)

func main() {
	// 多变量声明
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	// n1= 0 n2= 0 n3= 0

	var n4, name, n5 = 100, "tom", 888
	fmt.Println("n4=", n4, "name=", name, "n5=", n5)
	// n4= 100 name= tom n5= 888

	n6, sex, n7 := 111, "man", 777
	fmt.Println("n6=", n6, "sex=", sex, "n7=", n7)
	// n4= 100 name= tom n5= 888

	fmt.Println("n8=", n8, "name1=", name1, "n9=", n9)
	// n8= 800 name1= mary n9= 900
}

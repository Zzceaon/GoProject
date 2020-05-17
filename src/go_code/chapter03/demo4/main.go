package main
import "fmt"

func main() {
	// 该区域的数据值可以在同一类型范围内不断变化
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i=", i)
	// i = 1.2 不能改变数据类型
	// 只能在main范围内变化 且只能为int
}

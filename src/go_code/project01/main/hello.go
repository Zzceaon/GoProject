// 可go run 和 go build(实际生产中用go build再运行exe文件)
// exe文件在没有go的环境下任然可以运行,因此文件变大了(打包了go环境)
// go run需要go运行环境
// go build 可指定生成的文件名:go build -o myhello.exe hello.go
package main
import "fmt"
func main()  {
	// 快速复制:shift + alt + 上/下
	fmt.Println("hello world!")
	fmt.Println("hello world!")
	fmt.Println("hello world!")
}
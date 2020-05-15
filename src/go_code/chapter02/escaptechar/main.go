package main

import "fmt" // fmt包中提供格式化、输出、输入的函数
// package main 只能有一个main 不能有重名的
func main() {
	fmt.Println("tom\tjack")
	fmt.Println("tom\njack")
	fmt.Println("D:\\GoProject\\src\\go_code\\chapter02\\escaptechar")
	fmt.Println("tom\"jack\"")
	// \r 回车,从当前行的最前面开始输出,覆盖掉以前的内容
	fmt.Println("-------------\rjack")
	// 行注释
	// asfgasgasfa
	// safasfasfa

	// 块注释
	/*
	asfasfasf
	asfasfasfa
	asfasfasf
	*/
}
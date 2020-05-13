## Go特性
- 可go run 和 go build(实际生产中用go build再运行exe文件)
- exe文件在没有go的环境下任然可以运行,因此文件变大了(打包了go环境)
- go run需要go运行环境
- go build 可指定生成的文件名:go build -o myhello.exe hello.go
- 重命名 go build -o myhello.exe hello.go
- main是入口函数
- 一行只能写一条语句
- 定义的变量或者import的包如果没有用到,代码不能编译通过
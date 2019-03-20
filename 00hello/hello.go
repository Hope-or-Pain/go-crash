//package==project==workspace
package main

// 连接main包至fmt包（标准库）的所有代码、方法...
//同样可以连接的包还有math、debug、uploader等	Golang.org/pkg
import "fmt"

//主函数 main
func main() {
	fmt.Println("Hello, world.")
}
//go run编译并执行
//亦可go build编译后，执行./hello(.exe)
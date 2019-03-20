package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// // 01
// func main() {
// 	resp, err := http.Get("http://baidu.com")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(resp)
// }

// // 02
// func main() {
// 	resp, err := http.Get("http://google.com")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}
// 	// // 创建用于接收数据的[]byte字符串，99999位byte slice
// 	// bs := make([]byte, 99999)
// 	// // 将创建的slice传入read，read将response的内容存入传来的slice，故bs有了返回结果
// 	// // Reader interface内部，无论请求的response来的是啥，都转换输出存入[]byte
// 	// resp.Body.Read(bs)
// 	// fmt.Println(string(bs))

// 	// 以上三行代码可简写为
// 	// func Copy(dst Writer, src Reader) (written int64, err error)
// 	// 按住ctrl 点击Copy，会进入copy方法的源码,copyBuffer源码第390行左右
// 	// if buf == nil，通过make([]byte, size)创建用于接收数据的[]byte字符串，同上面的bs := make([]byte, 99999)
// 	io.Copy(os.Stdout, resp.Body)

// }

// 03
func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

// Reader、Writer常看io帮助
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}

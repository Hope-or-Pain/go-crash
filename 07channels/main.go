package main

import (
	"fmt"
	"net/http"
	"time"
)

// // 01
// func main() {
// 	links := []string{
// 		"http://google.com",
// 		"http://facebook.com",
// 		"http://amazon.com",
// 		"http://golang.org",
// 	}

// 	// // 非并发（行）情况下程序有等待，因为请求完第一个再请求第二个，效率低
// 	// for _, link := range links {
// 	// 	checkLink(link)
// 	// }

// 	// 非并发（行）情况下程序有等待，因为请求完第一个再请求第二个，效率低
// 	for _, link := range links {
// 		// 在调用函数时添加关键字go，创建新的go routine
// 		// 但如下程序不会有任何输出,因为main routine走完之后,下面循环产生的子routine被迫终结
// 		// 因此未来保证子routine正常运行，需要channel,Channel是不同routine间通信的唯一方式
// 		go checkLink(link)
// 	}
// }

// func checkLink(link string) {
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, "好像挂了")
// 		return
// 	}

// 	fmt.Println(link, "可访问")
// }

// // 02 channel
// func main() {
// 	links := []string{
// 		"http://google.com",
// 		"http://facebook.com",
// 		"http://amazon.com",
// 		"http://golang.org",
// 	}

// 	// 为routine通信创建channel
// 	c := make(chan string)

// 	// 非并发（行）情况下程序有等待，因为请求完第一个再请求第二个，效率低
// 	for _, link := range links {
// 		// 在调用函数时添加关键字go，创建新的go routine
// 		// 但如下程序不会有任何输出,因为main routine走完之后,下面循环产生的子routine被迫终结
// 		// 因此未来保证子routine正常运行，需要channel,Channel是不同routine间通信的唯一方式
// 		go checkLink(link, c)
// 	}

// 	// // 将channel得到的值输出，主routine程序会在此等待c接收到信息
// 	// // 但只有一行等待输出的c时，并行程序中最快的返回c值的网站响应后
// 	// // main routine发现自身没有需要再执行的代码，便推出了程序，channel只生效了一次
// 	// fmt.Println(<-c)
// 	// // 而再多加一行此代码时，生效两次，接收到上一个后
// 	// // main routine跳到此行代码并等待c接收值
// 	// // 以此类推，多加几次全都可输出，但多于循环次数时，程序会卡死在多余的等待上
// 	// fmt.Println(<-c)
// 	// fmt.Println(<-c)
// 	// fmt.Println(<-c)
// 	// // fmt.Println(<-c)

// 	// 故按照循环次数添加Println
// 	// 整个循环不会立即执行，而是像上面多行Println一样
// 	// 每到等待c的值时都会等待（阻塞调用）
// 	for i := 0; i < len(links); i++ {
// 		fmt.Println(<-c)
// 	}
// }

// // 接收string格式的channel参数
// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, "好像挂了")
// 		// 通过 <- 向channel发送信息
// 		c <- "网站挂了"
// 		return
// 	}

// 	fmt.Println(link, "可访问")
// 	// 通过 <- 向channel发送信息
// 	c <- "没问题可访问"
// }

// 02 channel
// 循环请求
func main() {
	links := []string{
		"http://google.com",
		"http://baidu.com",
		"http://amazon.com",
		"http://bing.com",
	}

	// 为routine通信创建channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// // 无限循环，一直获取这些网站
	// // 由于等待c值的返回，所以for循环李也是阻塞操作，不会一秒上百万次
	// for {
	// 	go checkLink(<-c, c)
	// }
	// 语义更明确的等效语法，其中l为link
	for l := range c {
		// 再者，虽然每次都有等待，但可再加上短暂的暂停
		// 但暂停语句在main routine时，收到google的返回channel信息，主程序便会冻结五秒
		// 这五秒中如果有其他子routine返回channel值也不会被接收，所以将sleep写入主程序是不正确的
		// time.Sleep(5 * time.Second)

		// go checkLink(l, c)

		// 使用function literal（函数文本，未命名的代码块）方式
		// 对应js中的Anonymous Function（匿名函数）: {}()
		// go func() {
		// 	time.Sleep(5 * time.Second)
		// 	checkLink(l, c)
		// }()
		// 以上程序代码会有警告，因为main routine中的l是变动的（循环 baidu到bing到...），
		// 而子程序处理时的预期只是传入时的值（如baidu）
		// 所以传入函数一个值（这个传入值会根据原值在内存中复制一份）可解决
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// 接收string格式的channel参数
func checkLink(link string, c chan string) {
	// sleep写在main toutine不对，但写在这里每次调用都等五秒再请求好像也不太对,如果有其他调用此方法不需等待的呢
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "好像挂了")
		// 通过 <- 向channel发送信息，link为string
		c <- link
		return
	}

	fmt.Println(link, "可访问")
	// 通过 <- 向channel发送信息
	c <- link
}

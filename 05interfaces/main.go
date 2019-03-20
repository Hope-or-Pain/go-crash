package main

import "fmt"

// 此bot interface为程序创建了名为bot的新type
// 此程序中所有 有getGreeting方法的 type，其元素都是bot的”荣誉成员“
// interface不同于concrate type(具体类型如int string)，不能直接生成bot类型的值
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// 通过interface将eb、sb作bot类型传入printGreeting()
	printGreeting(eb)
	printGreeting(sb)
}

// // 以下两方法用以输出英语和西语，但功能相同
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb englishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// 改用interface节约代码，此程序要求bot类型参数
// 传入eb或sb时,程序检查到二者皆是bot的“荣誉成员”(见上面bot定义处)故通过
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// 当 receiver的参数在函数体内没有被使用时，如下，可只写receiver的类型，去除eb
// func (eb englishBot) getGreeting() string {
func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"

}

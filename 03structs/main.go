package main

import "fmt"

// 定义名为person的struct
// 定义struct之前,需要定义所有可能的类型(int,string...)
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	xing string
	ming string
	// contact contactInfo,可简写为contactInfo
	// 即名为contactInfo的contactInfo类型
	contactInfo
}

// // 01
// func main() {
// 	// 值按照定义的顺序排列,易造成混淆
// 	wang := person{"王", "尼玛"}
// 	// 带上属性名会表述更明确
// 	zhang := person{xing: "张", ming: "三"}
// 	fmt.Println(wang, zhang)
// }

// // 02
// func main() {
// 	var wang person
// 	// golang中string空值"",int,float空值0,bool空之为false
// 	// 输出会展示为空,而不是undefined或null
// 	fmt.Println(wang)
// 	// 以下输出为{xing:,ming:}
// 	fmt.Printf("%+v", wang)
// }

// // 03
// func main() {
// 	var wang person
// 	wang.xing = "王"
// 	wang.ming = "尼玛"
// 	fmt.Println(wang)
// 	// 以下输出为{xing:,ming:}
// 	fmt.Printf("%+v", wang)
// }

// // 03
// func main() {
// 	zhang := person{
// 		xing: "张",
// 		ming: "三",
// 		contactInfo: contactInfo{
// 			email: "zhangsan@163.com",
// 			// 不同于js,定义多行struct时,即使最后一行也要有逗号,
// 			zipCode: 271000,
// 			// 下面一行也要有逗号,
// 		},
// 	}
// 	// updateName方法修改后,print输出仍为张三
// 	zhang.updateName("六")
// 	zhang.print()
// }

// // updateName方法修改后,上面print输出仍为张三
// // golang是值传递语言,不存在引用传递，所有传递至函数、receiver的值均在内存中复制，而不是传递内存地址
// // zhang调用下面方法传递zhang时,zhang的值从内存地址中完全复制,并作为p存入另一地址
// // 当修改p的值时,原内存地址上的原zhang的值不会被改变
// // 真正修改原zhang需要指针,见04
// func (p person) updateName(newMing string) {
// 	p.ming = newMing
// }

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }

// // 04
// func main() {
// 	zhang := person{
// 		xing: "张",
// 		ming: "三",
// 		contactInfo: contactInfo{
// 			email:   "zhangsan@163.com",
// 			zipCode: 271000,
// 		},
// 	}

// 	// &zhang获取zhang的内存地址
// 	// zhangPionter不是zhang的值,而是1000BH之类的内存地址的值

// 	// zhangPionter := &zhang
// 	// zhangPionter.updateName("六")

// 	// 以上两行代码可简写，golang语法糖允许这样写
// 	// zhang作为person类型，调用updateName（要求*person）传入方法时，
// 	// 会自动转换成指向person的指针
// 	zhang.updateName("六")

// 	zhang.print()
// }

// // *person位receiver为参数类型的描述
// // *person表示我们此方法只能被带指向person的指针调用,
// // 与下面*PointerToPerson的*作用不同
// func (PointerToPerson *person) updateName(newMing string) {
// 	// PointerToPerson即pointer(内存地址 &zhang)
// 	// *PointerToPerson is an operator
// 	// *PointerToPerson是指针PointerToPerson内存地址对应的值(zhang的值)
// 	(*PointerToPerson).ming = newMing
// }

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }

// 05
func main() {
	mySlice := []string{"1", "2", "3"}
	update(mySlice)
	fmt.Println(mySlice)
}

// 不同于例04，如下slice的值可以直接被改变而无需传指针
// 因为slice分为“slice文件”和“数组”两部分
// 第一部分存储 length（当前元素个数）、capacity（当前能储存多少元素）、pionter to head（指向数组头的指针）
// 第二部分是数组
// slice调用update传递值时，传递的是第一部分
// 方法内部的s的确是在内存复制后的，但复制的指针同样指向slice的第二部分
// 故改变s会改变mySlice的第二部分，也就是真正改变mySlice的值
func update(s []string) {
	s[0] = "0"
}

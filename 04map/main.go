package main

import "fmt"

// // 01
// func main() {
// 	// // 声明 map[key类型]value类型
// 不同于struct，struct是值传递结构，map为“引用传递”
// struct各元素类型随意；map的key须全为统一类型，value也是，但两者无需一致
// map的属性名可随时改变
// 真实场景中struct应用更多
// 	// colors := map[string]string{
// 	// 	"red":   "#ff0000",
// 	// 	"green": "#4bf745",
// 	// }

// 	// // 声明空map
// 	// var colors map[string]string

// 	// 亦可这样声明
// 	colors := make(map[int]string)
// 	colors[5] = "#ff0000"

// 	// 删除key
// 	delete(colors, 5)

// 	fmt.Println(colors)
// }

// 02
func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

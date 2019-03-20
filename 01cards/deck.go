package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//创建新 类型 ‘deck’
// which is a slice of string

// 类似于 继承 []string的属性
type deck []string

//(d deck)作为receiver
//d位为实际传入的变量（deck的副本）,可理解为deck的实例，在main.go中为cards
//可随意写作self、myDeck等，为节约资源，通常取type的首位（deck->d）
//deck位，为参数d的类型，即本go文件的“类”变量类型deck
//此方法为所有 ‘deck’类型的变量 添加print函数
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//此处不添加receiver，receiver只有在实例中要用到cards.xxx()方法时需要
// newDeck()后的deck为返回值的类型
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"tom", "cat", "dog"}
	cardValues := []string{"Ace", "Two", "Three"}
	//索引号在下面循环中使用不到，使用 _ 让程序忽略此位参数
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//参数d类型deck；参数handsize类型int
// ()后的deck为返回值的类型，此处返回根据handSize分割的两端Slice
// 此处不添加receiver是为避免语义不明,xxx.deal(5)给人感觉改变了原始变量xxx
func deal(d deck, handSize int) (deck, deck) {
	// return d[0:handSize],d[handSize:]
	// slice[:2]等同于slice[0:2]，为0位起到2位前,数学表示为[0，2）、{0,1}
	// slice[2:]为第2位至最后
	// retrn后逗号分割即可返回多个参数
	return d[:handSize], d[handSize:]
}

// 变量的“实例”重要用到xxx.toString()，故此处添加receiver
// 函数返回值，且返回类型为string，故此处添加string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// 变量的“实例”重要用到xxx.toString()，故此处添加receiver
// WriteFile(filename string, data []byte, perm os.FileMode) error
// 最后一位参数位权限配置，0666为所有人可更改
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// func ReadFile(filename string) ([]byte, error)
// 变量"实例"不会用到xxx.newDeckFromFile(),此处无receiver
func newDeckFromFile(filename string) deck {
	// bs,err分别接收Readfile方法的返回值byteSlice和error
	// 其中无错误是error返回nil
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 log the err and return a call to newDeck()
		// Option #2 log the err and entriely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// string.SSSSSplit，记得大写;string(bs)将[]byte转为[]string
	s := strings.Split(string(bs), ",")
	// 这里能够使用deck()是因为本程序开始设定deck基于string
	return deck(s)
}

// 类似于洗牌算法,将slice中某两位随即对调
func (d deck) shuffle() {
	for i := range d {
		// // 以下rand.Intn写法中,rand函数重复使用同一随机数表,导致每次执行的随机数结果一致
		// newPosition1 := rand.Intn(len(d) - 1)

		// 故通过时间创建随机数并改为以下写法
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		newPosition := r.Intn(len(d) - 1)
		// 错误写法: d已声明不能再":="而应是"="
		// // d[i], d[newPosition] := d[newPosition], d[i]
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

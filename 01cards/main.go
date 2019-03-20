package main

////01
// func main(){
// 	// // var card string = "Ace of Spades"
// 	// //亦可如下声明变量
// 	// code := "Ace of Spades"
// 	// //改变变量值时不再加“：”
// 	// code = "Something"

// 	card := newCard()

// 	fmt.Println(card)
// }
// func newCard() string {
// 	return "Five"
// }

////02
// func main() {
// 	//Slice是Array数组的一种，同Slice中元素必须为同一类型
// 	cards := []string{"Ace of Diamonds", newCard()}
// 	//append方法并不改变（cards，“added item”）中的cards
// 	// 此处cards = append（）的cards发生变化是因为赋值
// 	cards = append(cards, "Added Item")
// 	// range 循环；i index；:= 声明类似for(int i=1;...)
// 	for i, cards := range cards {
// 		fmt.Println(i, cards)
// 	}
// }

// //03
// func main() {
// 	//Slice是Array数组的一种，同Slice中元素必须为同一类型
// 	cards := deck{"Ace of Diamonds", newCard()}
// 	//append方法并不改变（cards，“added item”）中的cards
// 	// 此处cards = append（）的cards发生变化是因为赋值
// 	cards = append(cards, "Added Item")
// 	// range 循环；i index；:= 声明类似for(int i=1;...)
// 	cards.print()
// }

// //04
// func main() {
// 	cards := newDeck()
// 	cards.print()
// }

// // 05
// func main() {
// 	cards := newDeck()

// 	// 两个变量接收两个deck类型的返回值
// 	hand, remainingDeck := deal(cards, 5)
// 	hand.print()
// 	remainingDeck.print()
// }

// // 06
// func main() {
// 	word := "Hi there!"
// 	// []byte(string)方法将string转为byte（ASCII码）,[]byte类似[]string
// 	fmt.Println([]byte(word))

// 	cards := newDeck()
// 	fmt.Println(cards.toString())
// 	//deck转byte需要先转string
// 	fmt.Println([]byte(cards.toString()))
// }

// // 07
// func main() {
// 	cards := newDeck()
// 	cards.saveToFile("my_cards")
// }

// // 08
// func main() {
// 	cards := newDeckFromFile("my_cards")
// 	fmt.Println(cards)
// }

// 09
func main() {
	cards := newDeck()
	// cards.shuffle()
	cards.print()
}

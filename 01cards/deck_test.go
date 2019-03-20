package main

import (
	"os"
	"testing"
)

// 测试函数名首字符大写,类型处写*testing.T
// 此方法测试neweck()产生的deck是否符合要求,go test运行
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 9 {
		// Remain!	%v为输出控制符? 类似C中%s? 类似js中`${}`?
		t.Errorf("deck预期长度为9,测试得到 %v", len(d))
	}

	if d[0] != "Ace of tom" {
		t.Errorf("deck首位预期值为'Ace of tom',测试得到 %v", d[0])
	}

	if d[len(d)-1] != "Three of dog" {
		t.Errorf("deck尾位预期值为'Three of dog',测试得到 %v", d[len(d)-1])
	}
}

// 测试文件读写saveToFile和newDeckFromFile方法
// 用Test和Ande拼接长函数名是为了方便测试后用搜索查找
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 9 {
		t.Errorf("deck预期长度为9,测试得到 %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

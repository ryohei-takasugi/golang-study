package main

import (
	"fmt"
	"sample/error"
	"sample/fruit"
	"time"
)

func main() {
	// Fruit型の変数を作成
	var myFruit fruit.Fruit = fruit.Apple

	// toStringメソッドを使用して文字列を取得
	fmt.Println("Selected fruit:", myFruit)

	// 別のFruit定数を使用
	myFruit = fruit.Banana
	fmt.Println("Selected fruit:", myFruit)

	// 定義されていない値を使用
	myFruit = fruit.Fruit(3)
	fmt.Println("Selected fruit:", myFruit)

	myError := &error.MyError{
		When: time.Now(),
		What: "it didn't work",
	}
	fmt.Println(myError)
}

/*
	07_3. メソッド
	scriptの実行: $ go run chapter04/lesson12.go
*/
package main

import (
	"fmt"
)

/* 構造体にメソッドを定義する */
type user struct {
	name string
	age  int
}

/* メソッドは構造体の外で定義する */
func (u user) sayMessage(message string) { // 引数も戻り値も通常の関数のように書ける
	fmt.Println(message)
} // (u user) と関数名の前に書くことで, 構造体にメソッドを定義することができる

/* レシーバを使ったメソッドの定義 */
func (u user) selfIntroduce() {
	// uをレシーバと呼び, 構造体内のフィールドを参照できる
	fmt.Println("私の名前は" + u.name + "です。")
	fmt.Println(u.age, "歳独身で趣味はありません。")
}

func main() {
	/* メソッドを使ってみる */
	tanaka := user{name: "田中", age: 46}
	tanaka.sayMessage("こんにちは")
	tanaka.selfIntroduce()
}

/*
	02. 変数宣言
	scriptの実行: $ go run lesson02.go
*/
package main

import "fmt"

/* 変数定義（変更可能） */
var message string = "やあ、俺はゴウだ。よろしく。" // 型annotationあり（型は小文字始まり）

// 型推論も可能
var str = "ここにはstringが入るってわかってんだろ"

/* 複数の同時宣言 */
var foo, bar, buz string = "foo", "bar", "buz"
var n1, n2, n3 = 1, 2, 3

// 同じ型で複数宣言する場合
var (
	a string = "aaa"
	b        = "bbb"
	c        = "ccc"
	d int    = 123 // こいつだけ int
	e        = "eee"
)

/* 型の種類 */
var (
	word    string  // 文字列
	num1    int     // 整数
	num2    int8    // 8bit整数（-128〜127）
	boolean bool    // true/false
	float1  float32 // 小数点付き数値
	float2  float64 // 小数点付き数値
) // 他にもいろいろある

func main() {
	fmt.Println(message) // global で定義した message を使用
	printMessage()       // 1度目

	/* 関数内で再宣言可能*/
	message := "いや、俺はGo" // `:=` は関数内で再宣言するときに使用できる
	// var message string = "いや、俺はGo" // こう書いても良い

	fmt.Println(message) // 再宣言された message を使用
	printMessage()       // 2度目（こちらは変更されていない）

	/* 定数宣言（変更不可） */
	const name_c string = "progLearning"
	// name_c = "progLearning2" // const で宣言したので変更不可
	message = "いや、私がGoだ" // var で宣言したので変更可能

	/* 型の検証 */
	fmt.Printf("%T\n%T\n%T\n%T\n%T\n%T\n", word, num1, num2, boolean, float1, float2) // 空文字列

	/* 初期値に関して */
	fmt.Printf("%#v\n%#v\n%#v\n%#v\n%#v\n%#v\n", word, num1, num2, boolean, float1, float2) // それぞれ初期値が定義されている
}

func printMessage() {
	fmt.Println(message) // global で定義した message を常に使用
}

// 参考

// 数値の型
// https://qiita.com/tanaka0325/items/9c61a022cd32be0c65a6

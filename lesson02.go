package main

import "fmt"

/* 変数宣言 */

var message string = "やあ、俺はゴウ"

// 型annotation 型推論あり
// 文字列はダブルクォーテーションで囲む

/* 複数の宣言 */
var foo, bar, buz string = "foo", "bar", "buz"

// 同じ型で複数宣言する場合
var (
	a string = "aaa"
	b        = "bbb"
	c        = "ccc"
	d        = "ddd"
	e        = "eee"
)

/* 定数宣言 */
const name string = "progLearning"

/* string型 と int型 */
var word string
var num1 int
var num2 int8

// int8 ... 8bit整数（-128〜127）

func main() {
	fmt.Println(message)
	/* 関数内で再宣言（初期化） */
	message := "いや、俺はGo"
	fmt.Println(message)

	fmt.Println(word) // 空文字列
	fmt.Println(num)  // 0が初期値になっている
}

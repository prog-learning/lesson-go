/*
	02_1. 変数の宣言
	scriptの実行: $ go run chapter02_variable/lesson01go.go
*/
package main

import "fmt"

/* 宣言 */
var message1 string = "hi go" // 型も指定する
var message2 = "やあ、行け"        // 型を省略することも可能（型推論される）

/* 複数の同時宣言 */
var n1, n2, n3 int = 1, 2, 3            // 全ての型が同じ場合のみ
var foo, bar, buz = "foo", "bar", "buz" // 同様に型を省略可

// 以下のようにまとめることも可能
var (
	a string = "aaa"
	b        = "bbb"
	c        = "ccc"
	d int    = 123 // こいつだけ int型 にできる
	e        = "eee"
)

/* 以上に（main関数の外で）定義した変数はGlobal変数となる */
func main() {
	/* main関数内で宣言（このスコープでのみ使用できる変数） */
	// var message3 string = "Hello golang" // 変数が未使用の場合はエラー
	message3 := "Hello golang" // 略記法（型も省略して書く）

	/* 変数を出力 */
	fmt.Println(message3)

	message3 = "Hello golang" // 変数は再代入が可能

	/* 関数内ではGlobal変数を上書きして再宣言可能 */
	message1 := "HI GO"
	fmt.Println(message1)

	/* 変数の演算 */
	a, b := 10, 20
	fmt.Println(a + b) // 30
	// 複合代入演算子
	a += 5                        // a = a + 5
	b -= 5                        // b = b - 5
	fmt.Println("a:", a, "b:", b) // a: 15 b: 15
	// 単項演算子
	a++                           // a = a + 1
	b--                           // b = b - 1
	fmt.Println("a:", a, "b:", b) // a: 16 b: 14

	/* ブランク修飾子 */
	var _ = "ブランク修飾子" // 使用しない変数を宣言できる（エラーにならないのが利点）
	// _ := "ブランク修飾子"    // 省略形ではエラーになる

	/* 定数の宣言 */
	const name = "progLearning" // 使用しなくてもエラーにならない
	// name = "progLearning2" // 定数は変更できない
}

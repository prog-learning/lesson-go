/*
	07_1. 関数と引数
	scriptの実行: $ go run chapter04_function/lesson01.go
*/
package main

import (
	"fmt"
)

/* 宣言 */
func sayHello() {
	fmt.Println("関数で実行されました！")
}

// 実行は main関数内で行う
// 関数は（main）関数内で宣言できない

/* 引数 */
func sayMessage(message string) { // 引数の型は必須
	fmt.Println(message, "さん、ハロ〜！")
}

/* 複数の引数 */
func add1(x int, y int) {
	fmt.Println(x, "+", y, "=", x+y)
}

func add2(x, y int) { // 変数の定義と同様に省略可能
	fmt.Println(x, "+", y, "=", x+y)
}

func main() {
	/* それぞれの実行 */
	sayHello()

	sayMessage("まじ卍") // 引数を指定する

	add1(1, 2) // 2つの引数を指定する
	add2(4, 5)
}

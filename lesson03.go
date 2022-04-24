/*
	03. 演算子
	scriptの実行: $ go run lesson03.go
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	/* 文字列とは別に数字というものがある  */
	fmt.Println(12345) // "..."で囲まないと数字として扱われる

	// これらは演算子によって計算可能である
	fmt.Println(2 + 3)  // 5
	fmt.Println(2 - 3)  // -1
	fmt.Println(2 * 3)  // 6
	fmt.Println(10 / 5) // 2
	fmt.Println(10 % 3) // 1 余り

	/* べき乗はmathを使用する */
	fmt.Println(math.Pow(2, 3)) // 2の3乗で8
}

/* 補足 */

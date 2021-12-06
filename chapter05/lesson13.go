/*
	05_1. 条件演算子と論理演算子
	scriptの実行: $ go run chapter05/lesson13.go
*/
package main

import (
	"fmt"
)

func main() {
	/* 条件演算子... true か false が返る */
	const lang = "go"
	fmt.Println(lang == "go") // true
	fmt.Println(lang != "go") // false ! は否定

	score := 60
	fmt.Println(score > 80)  // false
	fmt.Println(score >= 40) // true
	fmt.Println(score < 60)  // false
	fmt.Println(score <= 60) // true

	/* 論理演算子 */
	// かつ（AND）
	fmt.Println(30 <= score && score <= 80) // true
	// または（OR）
	fmt.Println(lang == "go" || score > 80) // true
	// ではない（NOT）
	fmt.Println(!(score > 30)) // false
}

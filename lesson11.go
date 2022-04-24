/*
	11. 関数と戻り値
	scriptの実行: $ go run lesson11.go
*/
package main

import (
	"fmt"
	"math"
)

/* return で関数に戻り値を指定する */
func multi1(a, b int) int { // 戻り値の型を指定する（省略不可）
	return a * b
}

// 以下のように書くことともできる
func multi2(a, b int) int { // 戻り値の型を指定する（省略不可）
	c := a * b
	return c
}

/* 名前付き戻り値（戻り値となる変数を指定する） */
func multi3(a, b int) (c int) { // 戻り値となる変数 c を指定
	c = a * b // 変数を定義しなくても良くなる
	return c
}

// 名前付き戻り値を使うと、戻り値を省略できる
func multi4(a, b int) (c int) {
	c = a * b
	return // 戻り値の指定を省略可
}

/* 戻り値は複数にすることも可能 */
func getSquareRoots(x float64) (plus float64, minus float64) {
	plus = math.Sqrt(x) // 平方根を取得
	minus = -math.Sqrt(x)
	return plus, minus
	// return // この場合も省略可能
}

func main() {
	/* 戻り値を受け取る */
	result := multi1(2, 3)
	fmt.Println(result) // 6

	a, b, c := multi2(2, 3), multi3(2, 3), multi4(2, 3)
	fmt.Println(a, b, c) // 6 6 6

	/* 戻り値を複数受け取る */
	plus, minus := getSquareRoots(9)
	fmt.Println(plus, minus) // 3 -3
}

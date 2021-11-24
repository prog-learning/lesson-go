/*
	変数について
	実行： $ go run a-tour-of-go/section3.go
*/

package main

import (
	"fmt"
)

// var c, python, java bool // 初期値はfalseとなる これをゼロ値という
var c, python, java = true, false, "no!" // 初期値を指定する
// cobol := true // ここでこの記法は使えない

/* 定数 */
const Pi = 3.14
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	// ここで定義したものは, main関数内でのみ有効
	var i int // 初期値 0
	j := 3    // 短い書き方 型宣言を省略できる
	fmt.Println(i, j, c, python, java)

	/* 型の変換 */
	num := 42
	f := float64(num)
	u := uint(f)
	fmt.Println(num, f, u)

	/* 定数 */
	fmt.Println(Pi, Big*0.1, Small)
}

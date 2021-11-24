/*
	関数について
	実行： $ go run a-tour-of-go/section2.go
*/

package main

import (
	"fmt"
)

func add(x int, y int) int { // 引数と戻り値の型を指定する必要がある
	// func add(x, y int) int { // 省略形
	return x + y
}

func swap(x, y string) (string, string) { // 複数の戻り値
	return y, x
}

/* Named return values */
func split(sum int) (x, y int) { // 戻り値となる変数名を指定することができ,
	x = sum * 4 / 9
	y = sum - x
	// return x, y を省略して
	return // と書ける "naked return" という
	/* 長い処理がある関数では可読性が下がるので注意 */
}

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(27))
}

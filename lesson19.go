/*
	19. defer文
	scriptの実行: $ go run lesson19.go
*/
package main

import (
	"fmt"
)

/* 2つの処理を含む関数 */
func defer_test() {
	defer fmt.Println("deferをつかった処理") // 呼び出し元の関数の処理の最後に実行される

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

/* defer のスタック */
func defer_stack() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // 処理が蓄積されていく
	}
	fmt.Println("これから溜まった defer の処理を実行します")
	// 溜まった処理の順番に注目 First In Last Out となる
}

func main() {
	defer_test()
	defer_stack()
}

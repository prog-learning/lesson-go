/*
	01_2. Consoleに出力する
	scriptの実行: $ go run chapter01_basic/lesson02.go
*/
package main

import "fmt"

func main() {
	/* Println（プリントライン） で Console に出力  */
	fmt.Println("Hello world.")              // 文字列は "..." で囲む（ダブルクォーテーション○, シングルクォーテーション×）
	fmt.Println("hello", "world", "Println") // `,`で結合した際にはスペースが入る
}

/* 補足 */
// Println 以外にも Print, Printf などの複数の出力があるが,汎用性の高い Println をおさえておこう

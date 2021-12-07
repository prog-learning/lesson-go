/*
	07_1. pointer（ポインタ）
	scriptの実行: $ go run chapter07/lesson01.go
*/
package main

import "fmt"

/* ポインタとは ...宣言した変数などのメモリ領域のアドレス */
// 0x1040a0d0のような16進数で表される

func main() {
	/* ポインタの初期化 */
	var pointer *int             // *int...intのポインタ型
	fmt.Printf("%#x\n", pointer) // 0x0 // ポインタのゼロ値

	/* 変数からポインタを取得 */
	var lang string = "Go言語"        // まずは変数を宣言
	var langPointer *string = &lang // &lang...langのアドレスを取得
	fmt.Println(langPointer)        // 変数のアドレスを表示

	/* ポインタのアドレスに保存されている変数を取得 */
	fmt.Println(*langPointer)

	/* ポインタを通して書き換える */
	*langPointer = "Go language"
	fmt.Println(lang) // 変数が書き換わっている
}

// 参考

// https://qiita.com/Sekky0905/items/447efa04a95e3fec217f
// https://zenn.dev/uji/articles/f6ab9a06320294146733

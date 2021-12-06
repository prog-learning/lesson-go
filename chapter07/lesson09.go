/*
	09. pointer（ポインタ）
	scriptの実行: $ go run lesson09.go
*/
package main

import "fmt"

/* ポインタとは ...宣言した変数などのメモリ領域のアドレス */
// 0x1040a0d0のような16進数で表される

func main() {
	var lang string = "Go言語"  // 変数の宣言
	var langPointer = &lang   // langのアドレスを取得する
	fmt.Println(lang)         // 変数の値を表示
	fmt.Println(langPointer)  // 変数のアドレスを表示
	fmt.Println(*langPointer) // langPtrのアドレスにある値を表示 = lang
}

// 参考

// https://qiita.com/Sekky0905/items/447efa04a95e3fec217f
// https://zenn.dev/uji/articles/f6ab9a06320294146733

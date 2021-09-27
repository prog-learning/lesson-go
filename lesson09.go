/*
	09. ポインター
	scriptの実行: $ go run lesson09.go
*/
package main

import "fmt"

/* ポインターとは ...宣言した変数のメモリ領域のアドレスを確保する為の変数 */

func main() {
	var lang string = "Go言語"    // 変数の宣言
	var langPtr *string = &lang // langのアドレスを取得する
	fmt.Println(lang)           // 変数の値を表示
	fmt.Println(langPtr)        // 変数のアドレスを表示
	fmt.Println(*langPtr)       // langPtrのアドレスにある値を表示 = lang
}

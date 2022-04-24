/*
	20. pointer（ポインタ）
	scriptの実行: $ go run lesson20.go
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

	/* 実験 */
	var n int = 10 // int 10 を代入

	fmt.Println("n:", n)        // n: 10
	fmt.Printf("type: %T\n", n) // type: int

	fmt.Println("&n:", &n)       // &n: 0x14000014080 （アドレス）
	fmt.Printf("type: %T\n", &n) // type: *int （ポイント型）

	var p *int = &n // p に n のアドレスを代入
	// p := &n // でも可

	fmt.Println("p:", p)          // p: 0x14000014080
	fmt.Println("&p:", &p)        // p: 0x14000180020 （pのアドレス）
	fmt.Println("*p:", *p)        // *p: 10
	fmt.Println("*&n:", *&n)      // *&n: 10
	fmt.Printf("type: %T\n", p)   // type: *int
	fmt.Printf("type: %T\n", *p)  // type: int
	fmt.Printf("type: %T\n", *&n) // type: int

	*p = 99
	fmt.Println("n:", n) // n: 99 （nが変更される）

	/* 考察 */
	// &[変数]
	// アドレス演算子&を使って、変数のアドレスを取得できる
	// *[アドレスの入った変数]
	// アドレス演算子*を使って、アドレスに保存されている値を取得できる
	// pには、0x14000180020のアドレスにnのアドレス0x14000014080が値として格納されている状態
}

// 参考

// https://qiita.com/Sekky0905/items/447efa04a95e3fec217f
// https://zenn.dev/uji/articles/f6ab9a06320294146733
// https://code-graffiti.com/pointers-in-golang

/*
	01_1. 基礎知識
	scriptの実行: $ go run chapter01/lesson01.go
*/

/* package宣言層 */
package main // そのうち解説 とりあえずmainと書いておく

/* import層 */
import "fmt" // このファイルに置いて使用するpackageをimportする

/* global層 */
var langage = "golang" // globalスコープで変数宣言

/* main関数層 */
func main() {
	/* goでは常にmain()という関数が実行される */
	fmt.Println("hello" + langage) // console に hello golang!! と表示される
}

/* 必要そうな前提知識 */

// メモリを意識してコーディングする必要がある

// コード規約が厳しい（誰が書いても似たようなコードになる利点）

// Script言語ではなくコンパイル言語である

// フロントエンドと比較すると新しい概念が多い

// 静的型付け言語である

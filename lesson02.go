/*
	02_1. Consoleに出力する
	scriptの実行: $ go run lesson02.go
*/
package main

import "fmt"

func main() {
	/* Println（プリントライン） で Console に出力  */
	fmt.Println("Hello world.")              // 文字列は "..." で囲む（ダブルクォーテーション○, シングルクォーテーション×）
	fmt.Println("hello", "world", "Println") // `,`で結合した際にはスペースが入る
}

/* 補足 */
// Println 以外にも Print, Printf などの複数の出力があるが,汎用性の高い Println をまずはおさえておこう

/*
	02_2. Consoleに出力する
	scriptの実行: $ go run lesson01.go
*/

func main1() {
	/* Print ...出力 */
	fmt.Print("hello")
	fmt.Print("world.(改行されない)\n")         // `\n` をつけないと改行されない
	fmt.Print("hello", "world", "!!\n")   // `,`で結合できる
	fmt.Print("hello" + "world" + "!!\n") // `+`でも結合できる

	/* Println ...改行が含まれる 常用しそう 変数入れる場合は結合子 */
	fmt.Println("hello world. Println1")
	fmt.Println("hello", "world.", "Println2") // `,`で結合した際にはスペースが入る

	/* Sprint ...文字列を返す */
	hello1 := fmt.Sprint("Sprint", "Hello", "world!", "複数可能", "スペースが入らない")
	fmt.Println(hello1)
	hello2 := fmt.Sprintln("Sprintln", "Hello", "world!", "複数可能", "スペースが入る")
	fmt.Println(hello2)
	// `hello1 :=` は定数宣言の書き方（次回の説明で解説）

	/* printf ...フォーマットを使用 */
	myName := "Steve nobs"
	myAge := 200
	fmt.Printf("myName is %v.", myName)                                  // 変数を出力に組み込む
	fmt.Printf("myName is %v. %v が大好きです。%v 才\n", myName, "やきいも🍠", myAge) // %vで値が順番に入っていく感じ
	fmt.Printf("%#v\n %#v\n %#v\n", myName, "やきいも🍠", myAge)              // Goの文法でそのものを表示する
	fmt.Printf("%T\n %T\n %T\n", "文字列", 123, false)                      // 型を表示する
	fmt.Printf("単に%%を記号として出力したいとき")                                      // 記号の%を使いたいときは？
}

// 参考

// いろんなprint
// https://qiita.com/taji-taji/items/77845ef744da7c88a6fe

// printfのフォーマットの書き方
// https://qiita.com/rock619/items/14eb2b32f189514b5c3c

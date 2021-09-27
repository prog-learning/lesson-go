/*
	01. Consoleに出力する
	scriptの実行: $ go run lesson01.go
*/
package main

import "fmt"

// 文字列は"..."で囲む
// '...'は不可

func main() {
	/* Print ...出力 */
	fmt.Print("hello, ")
	fmt.Print("world.\n") // `\n` をつけないと改行されない

	/* Println ...改行が含まれる 常用しそう */
	fmt.Println("hello, Println1")
	fmt.Println("hello, Println2")

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
	fmt.Printf("myName is %v. %v が大好きです。%v 才\n", myName, "やきいも🍠", myAge) // %vで値が順番に入っていく
	fmt.Printf("%#v\n %#v\n %#v\n", myName, "やきいも🍠", myAge)              // Goの文法で表示する
	fmt.Printf("%T\n %T\n %T\n", "文字列", 123, false)                      // 型を表示する
	fmt.Printf("単に%%を記号として出力したいとき")                                      // 記号の%を使いたいとき
}

// 参考

// いろんなprint
// https://qiita.com/taji-taji/items/77845ef744da7c88a6fe

// printfのフォーマットの書き方
// https://qiita.com/rock619/items/14eb2b32f189514b5c3c

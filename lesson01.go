package main

import "fmt"

/* Hello World */

// 文字列は"..."で囲む
// '...'は不可

func main() {
	/* Print ...出力 */
	fmt.Print("hello, ")
	fmt.Print("world.\n") // 改行されない

	/* Println ...改行が含まれる */
	fmt.Println("hello, Println1")
	fmt.Println("hello, Println2")

	/* Sprint ...文字列を返す */
	hello := fmt.Sprintln("Hello", "world!")
	fmt.Println(hello)

	/* printf ...フォーマットを使用 */
	myName := "Steve nobs"
	myAge := 200
	fmt.Printf("myName of type is %T.\n", myName)
	fmt.Printf("myAge of type is %T.", myAge)

	/* 組み合わせてみる */
	introduce := fmt.Sprintf("My name is %v . %v years old", myName, myAge)
	fmt.Println(introduce)
}

// 参考
// https://qiita.com/taji-taji/items/77845ef744da7c88a6fe

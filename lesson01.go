/* Hello World */

package main

import "fmt"

// 文字列は"..."で囲まれている
// '...'は不可

func main() {
	/* Print ...出力 */
	fmt.Print("hello, ")
	fmt.Print("world.\n") // 改行されない

	/* Println ...改行が含まれる */
	fmt.Println("hello, Println1")
	fmt.Println("hello, Println2")

	/* Sprint ...フォーマット化された文字列を返す */
	hello := fmt.Sprintln("Hello", "world!")
	fmt.Print(hello)

	myName := "nobs"
	myAge := 200
	introduce := fmt.Sprintf("My name is %v . %v years old", myName, myAge)
	fmt.Print(introduce)
}

// 参考
// https://qiita.com/taji-taji/items/77845ef744da7c88a6fe

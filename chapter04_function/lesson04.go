/*
	07_4. 即時関数
	scriptの実行: $ go run chapter04_function/lesson04.go
*/
package main

import (
	"fmt"
)

func main() {
	/* 即時関数...宣言とともに実行する関数 */
	func() {
		fmt.Println("Hello, World!")
	}() // 関数内に書ける
}

/*
	15. module と package
	scriptの実行: $ go run lesson15.go
*/

package main

import (
	"fmt"

	"my_module/mypkg"
)

/* `go mod init モジュール名` をして, `go.mod` ファイルを作成する */

func main() {
	fmt.Println("Hello, World!")
	mypkg.Hello()
	mypkg.Hello2()
}

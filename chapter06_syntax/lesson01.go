/*
	06_1. go to文
	scriptの実行: $ go run chapter06_syntax/lesson01.go
*/
package main

import (
	"fmt"
)

func main() {
	/* goto文 */
	goto Here // 指定したタグがあれば、そこに飛ぶ
	fmt.Println("ぶおおおおおおおおおおおおん！🏍 💨 💨 💨 💨 💨 💨")
Here:
	fmt.Println("お休みなさい🛌")
}

/*
	05_3. go to文
	scriptの実行: $ go run chapter05/lesson15.go
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

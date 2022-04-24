/*
	21. 並行処理（goroutine と channel）
	scriptの実行: $ go run lesson21.go
*/
package main

import (
	"fmt"
	"time"
)

/* 時間のかかる関数を用意 */
func await() {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, "秒経過...")
	}
}
func awaitChannel(ch chan bool) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, "秒経過...")
	}
	ch <- true
}

func main() {
	/* 通常 */
	fmt.Println("開始")
	await() // この処理が終わってから次の処理
	fmt.Println("終了")

	/* goroutine */
	fmt.Println("開始")
	go await() // この処理を待たずに次の処理を実行する（これは待っているものがない状態）
	fmt.Println("終了")
	// awaitを待たずに次の処理を行ったため, awaitは実行されない

	/* channel */
	ch := make(chan bool)
	fmt.Println("開始")
	go awaitChannel(ch) // この処理を待たずに次の処理を実行する

	fmt.Println("またなくても平気だぜ")

	<-ch // awaitChannelの処理が終わるまで待つ
	fmt.Println("終わったよ")

}

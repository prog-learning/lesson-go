/*
	05_2. if文
	scriptの実行: $ go run chapter05/lesson14.go
*/
package main

import (
	"fmt"
)

func main() {
	/* if文 */
	ima := "asa"
	if ima == "asa" {
		fmt.Println("おはよう☀")
	}

	ima = "yoru"
	if ima == "asa" {
		fmt.Println("おはよう☀")
	}

	/* else文 */
	x, y := 1, 2
	if x > y {
		fmt.Println("xはyより大きい")
	} else {
		fmt.Println("xはyより大きくない🥺")
	}

	/* if else文 */
	score := 75
	if score > 80 {
		fmt.Println("最高！")
	} else if score > 60 {
		fmt.Println("いい感じ！")
	} else {
		fmt.Println("別にいんじゃね？")
	}

	/* 簡易変数定義 */
	if message, score := "おめでとう🎉", 98; score > 80 { // 条件演算子の前に;で区切って判定に使用する変数を定義できる
		// 定義した変数は, if文の中でのみ使用できる
		fmt.Println(message)
	}
}

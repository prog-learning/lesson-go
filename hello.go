package main

import (
	"fmt"
	"math/rand"
)

/* for文 */

func main() {
	for i := 0; i < 5; i++ { // 各項目省略可
		fmt.Println("うわ〜")
	}

	// while文風
	n := 0
	for n < 3 {
		fmt.Println("while文風で3回")
		n++
	}

	// break と continue
	var dice int
	for {
		dice = rand.Intn(6) + 1
		fmt.Println(dice)

		if dice == 1 {
			fmt.Println("おめでとうループを抜けます")
			break
		}
		if dice%2 == 0 {
			fmt.Println("ぐーすーだねー")
			continue
		}
	}
}

// 省略形などの書き方はない
// 誰が書いても同じになるのがGoの特徴

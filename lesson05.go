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

	/* 配列に繰り返し処理 */
	animals := [3]string{"cat", "dog", "pig"}

	// for文で配列の要素を取り出す
	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i])
	}

	// for文で配列を反復処理
	for index, value := range animals {
		fmt.Printf("%v番目の要素は%v\n", index, value)
	}

	// indexを省略可
	for _, value := range animals {
		fmt.Println(value)
	}
}

// 省略形などの書き方はない
// 誰が書いても同じになるのがGoの特徴

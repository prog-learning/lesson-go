/*
	17. for文（while文は存在しない）
	scriptの実行: $ go run lesson17.go
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/* 変数の定義; ループする条件; 処理後の演算など */
	for i := 0; i < 5; i++ { // 各項目省略可
		/* ループする処理 */
		fmt.Println("うわ〜", i)
	}

	/* while文風のfor文 */
	n := 0
	for n < 3 {
		fmt.Println("while文風で3回")
		n++
	}

	/* break と continue */
	var dice int
	for { // 全オプションを省略（永遠にループ）
		dice = rand.Intn(6) + 1 // 1~6のランダムな数字を生成
		fmt.Println("dice:", dice)

		if dice == 1 {
			fmt.Println("おめでとうループを抜けます")
			break // ループを抜ける
		}
		if dice%2 == 0 {
			fmt.Println("ぐーすーだねー。ふーん。")
			continue // ループ
		}
		/* 1が出るまで永遠にループ */
	}

	/* 配列に繰り返し処理 */
	animals := [3]string{"cat", "dog", "pig"}

	// for文で配列の要素を取り出す
	for i := 0; i < len(animals); i++ { // `len(array)`で配列arrayの要素数を取得
		fmt.Println(animals[i])
	}

	// for文で配列を反復処理
	for index, value := range animals { // `range array`で配列arrayのindexと要素を取得
		fmt.Printf("%v番目の要素は%v\n", index+1, value)
	}

	// for文でMapを反復処理
	profile := map[string]string{
		"name":    "gopher",
		"species": "gopher",
		"age":     "5",
	}
	for key, value := range profile {
		fmt.Printf("%v: %v\n", key, value)
	}

	// indexを省略する場合
	for _, value := range animals {
		fmt.Println(value)
	}
}

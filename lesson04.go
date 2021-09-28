/*
	04. 条件分岐 if文
	scriptの実行: $ go run lesson04.go
*/
package main

import (
	"fmt"
)

func main() {
	/* 条件演算子 */
	const lang = "go"
	fmt.Println(lang == "go") // true
	fmt.Println(lang != "go") // false

	score := 60
	fmt.Println(score > 60)  // true
	fmt.Println(score >= 60) // true
	fmt.Println(score < 60)  // false
	fmt.Println(score <= 60) // false

	/* 論理演算子 */
	// かつ（AND）
	fmt.Println(score <= 30 && 80 <= score) // false
	// または（OR）
	fmt.Println(lang == "go" || score > 80) // true
	// ではない（NOT）
	fmt.Println(!false && !(lang == "go")) // false

	/* if文 */
	x, y := 1, 2

	if x > y {
		fmt.Println("xはyより大きい")
	} else if x == y {
		fmt.Println("xとyは等しい")
	} else {
		fmt.Println("xはyより大きくない🥺")
	}

	/* goto文 */
	goto Here // 指定したタグがあれば、そこに飛ぶ
	fmt.Println("切ない感情")
Here:
	fmt.Println("嬉しい感情")

	/* ifとgotoを組み合わせてみる */
	status := "step2"
	if (status == "not started") || (status == "step1") {
		goto Initialized
	} else if status == "step2" {
		goto SecondPhase
	}
Initialized:
	fmt.Println("0.卵を割る")
	fmt.Println("1.卵を焼く")
SecondPhase:
	fmt.Println("2.ケチャップをかける")
	fmt.Println("3.完成")
}

// MEMO
// 省略形などの書き方はなく, 誰が書いても同じになるのがGoの特徴

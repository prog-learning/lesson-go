package main

import (
	"fmt"
)

/* if文 */

func main() {
	/* 条件演算子 */
	fmt.Println("GO" == "go") // false
	fmt.Println("GO" != "go") // true

	score := 60
	fmt.Println(score > 60)  // true
	fmt.Println(score >= 60) // true
	fmt.Println(score < 60)  // false
	fmt.Println(score <= 60) // false
	fmt.Println(score == 60) // false
	fmt.Println(score != 60) // true

	/* 論理演算子 */
	fmt.Println(score >= 60 && score <= 30) // false
	fmt.Println(score >= 60 || score <= 30) // true

	x, y := 1, 2

	if x < y {
		fmt.Println("xはyより大きい")
	} else if x == y {
		fmt.Println("xとyは等しい")
	} else {
		fmt.Println("xはyより大きくない🥺")
	}

}

// 省略形などの書き方はない
// 誰が書いても同じになるのがGoの特徴

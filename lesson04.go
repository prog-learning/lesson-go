package main

import (
	"fmt"
)

/* if文 */

func main() {
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

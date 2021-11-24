/*
	packageについて
	実行： $ go run a-tour-of-go/section1.go
*/

package main

import (
	"fmt"
	"math"      // mathパッケージをインポート
	"math/rand" // mathパッケージのrandパッケージをインポート
)

func main() {
	/* 一度目はランダムな整数が生成されるが, 二度目以降はその数字が固定される */
	x := rand.Intn(10)
	y := rand.Intn(99999)
	// z := math.rand.Intn(10) // mathパッケージのmath.rand.Intnは使えない
	fmt.Println(x, y)
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

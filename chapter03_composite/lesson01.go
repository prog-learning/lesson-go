/*
	03_1. 配列
	scriptの実行: $ go run chapter03_composite/lesson01.go
*/
package main

import "fmt"

func main() {
	/* 配列の定義 */
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	// `[5]int` ...int型で要素数が5の配列の型を指定（省略して型推論でもOK）
	// `[5]int{1, 2, 3, 4, 5}` ...int型の要素数が5の配列を作成
	//	※要素数を指定しなければならなく, あとから要素数を変更できない
	//	これはメモリの領域が固定されるということ

	fmt.Println(arr)        // 配列の要素を確認
	fmt.Printf("%T\n", arr) // 型を確認

	/* 次は文字列で */
	users := [...]string{"Tom", "Bill", "Mike", "John", "Mary"}
	// 要素数が未定の場合`[...]`と書くこともできる
	// ※要素数が動的になるわけではなく, あくまでも生成されたときに要素の個数が決まるということ

	/* 配列の呼び出し */
	fmt.Println(users[2] + " is a good") // Mike is a good 3番目(index2)の要素を取得

	/* 配列の要素を変更する */
	users[2] = "Mickey"
	fmt.Println(users[2]) // Mickey

	/* 配列の個数を取得する */
	fmt.Println(len(users)) // 5

	/* 配列の一部を切り取る（一部の参照データ） */
	group1 := users[2:4]           // 配列の3番目(index2)から4番目まで(index4を含めない)を切り取る
	group2 := users[:3]            // 配列の最初からから3番目(index3を含めない)までを切り取る
	group3 := users[2:]            // 配列の3番目(index2)から4最後までを切り取る
	fmt.Println("group1:", group1) // group1: [Mickey John]
	fmt.Println("group2:", group2) // group2: [Tom Bill Mike]
	fmt.Println("group3:", group3) // group3: [John Mary]
	fmt.Println(users)             // [Tom Bill Mickey John Mary] これは配列の要素を変更していない(非破壊的)

	group1[1] = "Mickey" // 切り取った配列の要素を変更
	fmt.Println(users)   // [Tom Bill Mickey Mickey John Mary] もとの配列に影響されるので注意
}

/* 補足 */
// 配列では個数を操作することができない
/*
	略記
	以下同義
	a[0:10]
	a[:10]
	a[0:]
	a[:]
*/

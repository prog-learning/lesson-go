/*
	03. 配列
	scriptの実行: $ go run lesson03.go
*/
package main

import "fmt"

func main() {
	/* 配列の定義 */
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	/*
		POINT:
		`[5]int` ...int型で要素数が5の配列の型を指定（省略して型推論でもOK）
		`[5]int{1, 2, 3, 4, 5}` ...int型の要素数が5の配列を作成
		※要素数を指定しなければならなく, あとから要素数を変更できない
	*/
	fmt.Printf("%T\n", arr) // 型を確認
	fmt.Println(arr)

	/* 次は文字列で */
	users := [...]string{"Tom", "Bill", "Mike", "John"}
	// 要素数が未定の場合`[...]`と書く（要素数が動的になるわけではないので注意）
	fmt.Println(users[2] + " is a good") // Mike is a good 3番目(index2)の要素を取得
	fmt.Println(users[0] == "Tom")       // true
	fmt.Printf("%T\n", users)            // 型を確認

	/* 配列の要素を変更する */
	users[2] = "Mickey"
	fmt.Println(users[2]) // Mickey

	/* 配列の要素数を動的（変更可能）にする */
	animals := []string{"cat", "dog", "pig", "cow", "horse"}
	fmt.Printf("%T\n", animals) // 型を確認
	/* 配列に追加 */
	animals = append(animals, "duck") // 再代入している
	fmt.Println("animals:", animals)  // [cat dog pig cow horse duck]

	/* 配列の一部を切り取る（一部の参照データ） */
	cutAnimals1 := animals[2:4]              // 配列の3番目(index2)から4番目まで(index4を含めない)を切り取る
	cutAnimals2 := animals[:3]               // 配列の最初からから3番目(index3を含めない)までを切り取る
	cutAnimals3 := animals[2:]               // 配列の3番目(index2)から4最後までを切り取る
	fmt.Println("cutAnimals1:", cutAnimals1) // [pig cow]
	fmt.Println("cutAnimals2:", cutAnimals2) // [cat dog pig]
	fmt.Println("cutAnimals3:", cutAnimals3) // [pig cow horse duck]
	fmt.Println(animals)                     // [cat dog pig cow horse duck] もとの配列は変更されない（非破壊的）
	cutAnimals1[0] = "mouse"                 // 切り取った配列の要素を変更
	fmt.Println(animals)                     // [cat dog mouse cow horse duck] もとの配列に影響されるので注意

	/* 配列の要素を削除する */
	animals = append(animals[:2], animals[4:]...) // 最初から2番目まで(2+1を含めない)に、(4+1)番目から最後までを追加
	fmt.Println(animals)                          // [cat dog horse duck]
}

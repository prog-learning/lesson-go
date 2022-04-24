/*
	07. Slice（スライス）
	scriptの実行: $ go run lesson07.go
*/
package main

import "fmt"

// スライスは配列と異なり, 可変長であり, 要素の個数を変更できる
// 基本的な操作は配列と同じ

func main() {
	/* 定義 */
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice) // [1 2 3 4 5]

	/* 参照 */
	fmt.Println(slice[1]) // 2

	/* 長さ */
	fmt.Println(len(slice)) // 5

	/* 追加 */
	slice = append(slice, 6)       // appendは新しいスライスを返す
	fmt.Println(slice)             // [1 2 3 4 5 6]
	slice = append(slice, 7, 8, 9) // 複数同時に追加できる
	fmt.Println(slice)             // [1 2 3 4 5 6 7 8 9]

	/* 削除 */
	slice = append(slice[:3], slice[6:]...) // 0~3番目を削除し, 6~最後までを追加
	fmt.Println(slice)                      // [1 2 3 6 7 8 9]

	/* スライスのコピー */
	slice2 := make([]int, len(slice)) // len(slice)はスライスの要素数
	copy(slice2, slice)               // copyはコピーする
	fmt.Println(slice2)               // [1 2 3 6 7 8 9]
}

/*
	08. Map（マップ）
	scriptの実行: $ go run lesson08.go
*/
package main

import "fmt"

func main() {
	/* 定義 */
	var item map[string]int = map[string]int{ // `map[string]int` ... keyが string で値が int の map
		"no":         12345,
		"price":      5000,
		"numOfItems": 99,
	} // すべての型が int に統一されることになる

	/* 参照 */
	fmt.Println(item["price"]) // 5000

	/* 要素の追加 */
	item["serial"] = 1234567890
	fmt.Println(item) // map[no:12345 price:5000 numOfItems:99 serial:1234567890]

	/* 要素の削除 */
	delete(item, "numOfItems")
	fmt.Println(item) // map[no:12345 price:5000 serial:1234567890]

	/* keyの存在を調べる */
	value, exist := item["price"] // 代入で2つ目の値を受け取ることができる
	fmt.Println(value, exist)     // 5000 true

}

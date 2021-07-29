package main

import "fmt"

/* 配列 */

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	// [5]int ...int型で要素数が5の配列の型
	// [5]int{1, 2, 3, 4, 5} ...int型で要素数が5の配列を作成
	fmt.Println(arr)

	users := [5]string{"Tom", "Bill", "Mike", "John", "Sally"}
	fmt.Println(users[2] + " is a good") // Mike is a good
	fmt.Println(users[0] == "Tom")       // true

	/* 配列の要素を変更する */
	users[2] = "Mikey"
	fmt.Println(users[2]) // Mikey

	/* 配列の要素を追加する */
	animals := []string{"cat", "dog", "pig", "cow", "horse"}
	animals = append(animals, "duck")
	fmt.Println("animals:", animals) // [cat dog pig cow horse duck]

	/* 配列の一部を切り取る */
	cutAnimals1 := animals[2:4]              // 配列の(2+1)番目から4番目まで(4+1を含めない)を切り取る
	cutAnimals2 := animals[:3]               // 配列の(2+1)番目から4番目まで(4+1を含めない)を切り取る
	cutAnimals3 := animals[2:]               // 配列の(2+1)番目から4番目まで(4+1を含めない)を切り取る
	fmt.Println("cutAnimals1:", cutAnimals1) // [pig cow]
	fmt.Println("cutAnimals2:", cutAnimals2) // [cat dog pig]
	fmt.Println("cutAnimals3:", cutAnimals3) // [pig cow horse duck]
	fmt.Println(animals)                     // [cat dog pig cow horse duck] 非破壊的

	/* 配列の要素を削除する */
	animals = append(animals[:3], animals[4:]...)
	fmt.Println(animals) // [cat dog pig horse duck]

}

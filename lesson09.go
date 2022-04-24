/*
	09. Struct（構造体）
	scriptの実行: $ go run lesson09.go
*/
package main

import (
	"fmt"
)

func main() {
	/* 定義 */
	type user struct {
		/* フィールド */
		name                string
		age                 int
		login               bool
		firstName, lastName string // 型が同じ場合はこのように省略できる
	}

	/* 構造体の初期化 */
	var user0 user
	fmt.Println(user0) // フィールドにはそれぞれのゼロ値が入っている

	/* フィールドを指定しながら初期化 */
	user1 := user{"pokemon-get", 7, true, "okido", "hakase"}
	// 順番に指定することでフィールドを指定できるが, すべて順番通りに指定する必要がある

	/* 指定したいものだけ指定しながら初期化 */
	user2 := user{
		login: true,
		name:  "satoshino-pikachu",
		age:   18,
	} // この形式で書けば,必要箇所だけ指定できるし,順番も問わない

	fmt.Println(user1) // {pokemon get 7 true okido hakase}
	fmt.Println(user2) // {true satoshino pikachu 18} // 指定した順番になる

	/* 書き換える（再代入） */
	user2.name = "takesino-iwaaku"
	fmt.Println(user2)

	/* 参照 */
	fmt.Println(user2.name, "is login.") // takesino-iwaaku is login.
	// マップと構造体は異なるので注意

	/* 構造体の埋め込み */
	type birthday1 struct {
		year  int
		month int
	}
	type birthday2 struct {
		year  int
		month int
	}
	// 上の2つの構造体を埋め込んで定義
	type profile struct {
		birthday1           // 埋め込み（読み書きの省略が可能）
		birthday2 birthday2 // 名前付きで埋め込み（読み書きの省略はできない）
		email     string
		bio       string
	}

	// profile を使い2つを検証していく
	profile0 := profile{} // 初期化
	// birthday1
	profile0.birthday1.year = 1
	profile0.month = 2    // .birthday1 を省略できる
	fmt.Println(profile0) // {{1 2} {0 0}  }

	// birthday2
	profile0.birthday2.year = 3 // 省略した書き方はない
	profile0.month = 4          // .birthday1.month に書き込まれる
	fmt.Println(profile0)       // {{1 4} {3 0}  }

	// 代入しながら初期化するときは構造体を意識した書き方（フラットにできない）
	profile1 := profile{
		birthday1: birthday1{
			year:  1,
			month: 2,
		},
		birthday2: birthday2{
			year:  3,
			month: 4,
		},
		email: "happy-birthday@gmail.com",
		bio:   "congratulations!!",
	}
	fmt.Println(profile1) // {{1 2} {3 4} happy-birthday@gmail.com congratulations!!}
}

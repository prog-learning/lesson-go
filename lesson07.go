/*
	07. map（JavaScriptで言うところのオブジェクト）
	scriptの実行: $ go run lesson07.go
*/
package main

import (
	"fmt"
)

func main() {
	/* 定義 */
	user := map[string]string{ // `map[string]string` ...keyがstringで値がstringのmap
		"name": "ピエン 翼",
		"age":  "7",
	}
	fmt.Println(user["name"]) // 呼び出し

	/* 要素の追加 */
	user["tell"] = "090-1234-5678"
	fmt.Println(user)

	/* 要素の削除 */
	delete(user, "age")
	fmt.Println(user)

	/* keyの値と有無を調べる */
	if value, ok := user["age"]; ok {
		fmt.Printf("age is exist.value is %v\n", value)
	} else {
		fmt.Println("age is not exist.")
	}

	/* mapの型を定義 */
	type userType struct {
		name  string
		age   int
		login bool
	}

	user1 := userType{
		name:  "ピエン 翼",
		age:   7,
		login: true,
	}

	fmt.Println(user1.name, "is login.")
}

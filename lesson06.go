/*
	06. 関数
	scriptの実行: $ go run lesson06.go
*/
package main

import (
	"fmt"
	"strings"
)

/* 定義 */
func firstFunc() {
	fmt.Println("初めての関数で実行されています")
}

func sayHello(name string) { // 引数の型は必須
	fmt.Println(name, "さん、ハロ〜！")
}

func allEat(foods []string) {
	for _, food := range foods {
		fmt.Printf("%vを食べる \n", food)
	}
}

/* 戻り値 return */
func nameLength(name string) int { // 戻り値の型は必須
	length := len(name)
	return length
}

// 名前付き戻り値
func greet(name string) (message string) { // 戻り値を指定しておく
	message = name + "さん、こんにちは"
	return
	// return message と書かなくても良い
}

// 複数の戻り値を返す
func getInitial(name string) (string, string) {
	var initial []string
	names := strings.Split(name, " ")
	for _, name := range names {
		initial = append(initial, strings.ToUpper(name[:1]))
	}

	if len(initial) > 1 {
		return initial[0], initial[1]
	}
	return initial[0], ""
}

func main() {
	firstFunc()
	sayHello("pien🥺")

	foods := []string{"ラーメン", "カレーライス", "焼肉定食"}
	allEat(foods)

	length := nameLength("pien🥺")
	fmt.Println(length)

	message := greet("pien🥺")
	fmt.Println(message)

	initial1, initial2 := getInitial("pien pien")
	fmt.Printf("%v.%v\n", initial1, initial2)

}

package main

import (
	"fmt"
	"strings"
)

/* 関数 */

func greet(name string) {
	fmt.Println(name, "さん、こんにちは")
}

func allEat(foods []string) {
	for _, food := range foods {
		fmt.Printf("%vを食べる \n", food)
	}
}

/* 戻り値 */

func getInitial(name string) (string, string) { // 戻り値の型を指定する
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
	greet("pien")

	foods := []string{"ラーメン", "カレーライス", "焼肉定食"}
	allEat(foods)

	initial1, initial2 := getInitial("pien pien")
	fmt.Printf("%v.%v", initial1, initial2)

}

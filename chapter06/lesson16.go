/*
	06_1. Switch文
	scriptの実行: $ go run chapter06/lesson16.go
*/
package main

import "fmt"

func main() {
	signal := "blue"

	switch signal {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("caution")
	case "green", "blue": // 複数の case を指定する場合はカンマで区切る
		fmt.Println("Go")
	default:
		fmt.Println("wrong")
	}
}

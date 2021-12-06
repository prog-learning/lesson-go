package mypkg

import "fmt"

/* 外部でしようされるものは大文字始まりで */
func Hello2() {
	fmt.Println("hello package by file2.go")
}

package mypkg // パッケージ名 = ディレクトリ名とすること

import "fmt"

/* 外部でしようされるものは大文字始まりで */
func Hello() {
	fmt.Println("hello package")
}

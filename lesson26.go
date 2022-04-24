/*
20. pointer（ポインタ）を使った書き換え
scriptの実行: $ go run lesson26.go
*/
package main

import "fmt"

func putByValue(n int) {
	n = 99
	fmt.Println("&v:", &v) // 0x100e2c368
	fmt.Println("&n:", &n) // 0x14000014080
	fmt.Println("In putByValue v:", v)
}
func putByPointer(p *int) {
	*p = 99
	fmt.Println("In putByPointer v:", v) // v: 10
}

var v = 10

func main() {
	fmt.Println("=initial==============")
	fmt.Println("v:", v)   // v: 10
	fmt.Println("&v:", &v) // &v: 0x14000014080 （アドレス）

	fmt.Println("======================")
	putByValue(v)
	fmt.Println("=After putByValue=====")
	fmt.Println("v:", v) // v: 10

	fmt.Println("======================")
	putByPointer(&v)
	fmt.Println("=After putByPointer====")
	fmt.Println("v:", v) // v: 99

}

// 考察
// 関数内で変数を変更する場合は、その変数のポインタアドレスを渡す

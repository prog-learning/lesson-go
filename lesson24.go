package main

/*
	24. Interface
	scriptの実行: $ go run lesson24.go
*/

import (
	"fmt"
	"strconv"
)

type Calculate interface {
	Increment() int
	Decrement() int
	Double() int
	Multiply(int) int
	Reset() int
}

type Stringer struct {
	value string
}

type Integer struct {
	value int
}

/* Str用の関数群 */
func (s Stringer) Increment() int {
	fmt.Printf("type: %T\n", s.value)
	/* Intに変換 */
	i, err := strconv.Atoi(s.value)
	if err != nil {
		fmt.Println(err)
	}
	/* Strに変換 */
	return i + 1
}
func (s Stringer) Decrement() int {
	fmt.Printf("type: %T\n", s.value)
	/* Intに変換 */
	i, err := strconv.Atoi(s.value)
	if err != nil {
		fmt.Println(err)
	}
	return i - 1
}
func (s Stringer) Double() int {
	/* Intに変換 */
	i, err := strconv.Atoi(s.value)
	if err != nil {
		fmt.Println(err)
	}
	return i * 2
}
func (s Stringer) Multiply(n int) int {
	/* Intに変換 */
	i, err := strconv.Atoi(s.value)
	if err != nil {
		fmt.Println(err)
	}
	return i * n
}
func (s Stringer) Reset() int {
	return 0
}

/* Int用の関数群 */
func (s Integer) Increment() int {
	fmt.Printf("type: %T\n", s.value)
	return s.value + 1
}
func (s Integer) Decrement() int {
	fmt.Printf("type: %T\n", s.value)
	return s.value - 1
}
func (s Integer) Double() int {
	return s.value * 2
}
func (s Integer) Multiply(n int) int {
	return s.value * n
}
func (s Integer) Reset() int {
	return 0
}

/* Interfaceを引数にした汎用性の高い関数を作成 */
func CalculateAll(c Calculate, i int) int {
	return c.Increment()
	return c.Decrement()
	return c.Double()
	return c.Multiply(i)
	return c.Reset()
}

func main() {
	var s Stringer = Stringer{value: "100"}
	var i Integer = Integer{value: 999}
	var calc Calculate
	calc = s
	fmt.Println(calc.Increment())
	fmt.Println(calc.Decrement())
	fmt.Println(calc.Double())
	fmt.Println(calc.Multiply(5))

	calc = i
	fmt.Println(calc.Increment())
	fmt.Println(calc.Decrement())
	fmt.Println(calc.Double())
	fmt.Println(calc.Multiply(5))

}

// 参考

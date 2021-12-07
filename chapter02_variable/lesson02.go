/*
	02_2. 変数の初期化と型について
	scriptの実行: $ go run chapter02_variable/lesson02.go
*/
package main

import "fmt"

func main() {

	/* 変数の初期化 */
	var message string   // 代入するものを書かない（型は必ず書く）
	fmt.Println(message) // 空文字列
	// 代入をせずに宣言することを初期化といい, それぞれにはゼロ値（初期値）が定義される

	/* それぞれの型のゼロ値 */
	var s string
	var i int
	var b bool
	fmt.Println(s, i, b) // 空文字列, 0, false
	// 他にも型が複数ある

	/* 様々な型の種類 */
	var (
		string1  string  // 文字列
		uint1    uint    // 符号なし整数
		uint2    uint8   // 8bit 符号なし整数
		uint3    uint16  // 16bit 符号なし整数 // 64bitまである
		int1     int     // 符号付き整数（結構おっきくてもいける）
		int2     int8    // 8bit符号付き整数（-128〜127）
		int3     int16   // 16bit符号付き整数（-32768〜32767） // 64bitまである
		float1   float32 // 32bit浮動小数点数（小数点も含めそこそこでかい数字を扱える）
		float2   float64 //  64bit浮動小数点数(doubleとも言われる)
		boolean1 bool    // true / false
	) // データはbit容量で管理されるのでこのように数字には複数の型が存在する
	fmt.Println(string1, uint1, uint2, uint3, int1, int2, int3, boolean1, float1, float2) // 空文字列, 0, 0, 0, 0, 0, 0, false, 0, 0

	/* 型の確認 */
	fmt.Printf("string1  ... 型: %T, ゼロ値: %v\n", string1, string1)
	fmt.Printf("uint1    ... 型: %T, ゼロ値: %v\n", uint1, uint1)
	fmt.Printf("uint2    ... 型: %T, ゼロ値: %v\n", uint2, uint2)
	fmt.Printf("uint3    ... 型: %T, ゼロ値: %v\n", uint3, uint3)
	fmt.Printf("int1     ... 型: %T, ゼロ値: %v\n", int1, int1)
	fmt.Printf("int2     ... 型: %T, ゼロ値: %v\n", int2, int2)
	fmt.Printf("int3     ... 型: %T, ゼロ値: %v\n", int3, int3)
	fmt.Printf("float1   ... 型: %T, ゼロ値: %v\n", float1, float1)
	fmt.Printf("float2   ... 型: %T, ゼロ値: %v\n", float2, float2)
	fmt.Printf("boolean1 ... 型: %T, ゼロ値: %v\n", boolean1, boolean1)
	// fmt.Printf()内で%Tと書くと型を表示できる

	/* 型の変換 */
	int1 = 123
	// int2 = int1 // 型の不一致でエラー
	int2 = int8(int1)
	fmt.Printf("int2     ... 型: %T, 値: %v\n", int2, int2)

}

// 参考

// 数値の型
// https://qiita.com/tanaka0325/items/9c61a022cd32be0c65a6

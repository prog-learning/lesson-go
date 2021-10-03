package main

/*
	10. サーバーを起動する
	scriptの実行: $ go run lesson10.go
*/

import (
	"fmt"
	"net/http"
)

/* Responseに使用する関数 */
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Golang") // 書き込み先をwに指定する書き方（https://pkg.go.dev/fmt#Fprintf）
}

type Hello string

// Helloに ServeHTTP 関数を追加
/* この書き方なに！？ */
func (value Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, value)
}

func main() {
	/* エンドポイントの作成 */
	http.HandleFunc("/", hello) // (Endpoint, ハンドリングする処理)を指定する
	// 無名関数を使用
	http.HandleFunc("/lang", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Golang")
	})
	// 第2引数に引数を使いたい場合
	http.Handle("/big", Hello("HELLO"))

	/* サーバーの起動（最後に書く） */
	fmt.Println("lesten http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

// 参考
// https://www.yoheim.net/blog.php?q=20170403
// https://qiita.com/entaku0818/items/c29add790718c215381e

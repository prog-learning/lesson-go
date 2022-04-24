# Lesson Go

Go 言語を学ぶ

各ファイルで main 関数を定義しているので,重複していてエラーが出ますが,大切なことは気にしないことです.

## References

- https://golang.org/
- https://qiita.com/tenntenn/items/0e33a4959250d1a55045
- https://gihyo.jp/dev/feature/01/go_4beginners
- https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/
- https://qiita.com/tfrcm/items/e2a3d7ce7ab8868e37f7

## Installation

公式 https://golang.org/

brew を使う場合

```sh
brew install go
```

## Setting

Go は Prettier がきかないので,Go の拡張機能を install し,以下を VSCode の setting.json に追加した

```json
"[go]": {
  "editor.defaultFormatter": "golang.go"
},
```

これで自動整形ができるようになった.

## Go command

実行

```sh
go run ファイル名
```

コード整形

```sh
go fmt ファイル名
```

## Lessons

1. 基礎知識
2. Console に出力する
3. 演算子
4. 変数の宣言

- 基本的な知識
- コンソールに表示する
- 変数
- 配列
- スライス
- Map マップ
- Struct（構造体）
- if 文
- switch 文
- for 文
- defer 文
- goto 文
- 関数
- レシーバと構造体メソッド
- 即時関数
- ポインタ
- ポインタを使った構造体
- パッケージ
  - fmt
  - math
  - strings
- 並行処理（channel と goroutine）
- net/http

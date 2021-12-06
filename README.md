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

1. Hello world と print
2. 変数宣言
3. 配列
4. if 文
5. for 文
6. 関数
7. map (オブジェクト)
8. switch 文
9. pointer
10. http

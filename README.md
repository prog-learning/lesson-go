# Lesson Go

Go 言語を学ぶ

## Reference

https://qiita.com/tenntenn/items/0e33a4959250d1a55045
https://gihyo.jp/dev/feature/01/go_4beginners
https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/

## Installation

公式 https://golang.org/

brew

```
brew install go
```

## Setting

Go は Prettier がきかないので,Go の拡張機能を install し,以下を VSCode の setting.json に追加

```
"[go]": {
  "editor.defaultFormatter": "golang.go"
},
```

## 基本知識

実行

```
go run ファイル名
```

コード整形

```
go fmt ファイル名
```

## Lessons

- Lesson01
  Hello world
- Lesson02

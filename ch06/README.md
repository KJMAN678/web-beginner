## ListenAndServe
```sh
# Linux コマンドでローカルサーバー立ち上げ
echo 'Hello, Web application!' | nc -l 8080

# go でローカルサーバー立ち上げ
go run ch06/ListenAndServe/main.go
```
# ローカルサーバーのURL
http://localhost:8080/

A に B を書き込み
fmt.Fprint(A, B)

### URLの設定
```sh
http.HandleFunc("/", func)
```

### ローカルサーバーで8080ポート立ち上げ

立ち上がれば nil を返す
```sh
http.ListenAndServe(":8080", nil)
```

### セイウチ演算子
```sh
var err = http.ListenAndServe(":8080", nil)
err := http.ListenAndServe(":8080", nil)
```

### エラーハンドリング

```sh
err := func()
if err != nil {
    log.Fatal("failed to start:", err)
}
```

## FileServer static/xxx.html を読み込んでローカルサーバー立ち上げ

```sh
go run ch06/FileServer/main.go
```
# ローカルサーバーのURL
http://localhost:8080/hello.html

### 指定のフォルダ以下のHTMLファイルをローカルサーバーで立上げ

```sh
http.FileServer(http.Dir("hoge-dir")))
```

### CSSとJavaScriptも読み込みたい
```sh
go run ch06/FileServerCSSJavaScript/main.go
```

- static フォルダ以下に css と JavaScript も追加。
- CSS: h1 の色変更
- JavaScript: ButtonClick による文字列の追加

## TodoList
```sh
go run ch06/TodoList/main.go
```
http://localhost:8080/todo

- http.StripPrefix(): http.FileServer() に渡すパスを調整している。static という文字列を除いている
- template.ParseFIle() ... template ディレクトリに用意した html ファイルを読み込み、テンプレートとして解析する

## TodoList を追加
```sh
go run ch06/TodoListAdd/main.go
```
http://localhost:8080/todo
http://localhost:8080/add

- /add パスにアクセスすると、handleAdd()を実行する
```sh
http.HandleFunc("/add", handleAdd)
```

- post メソッドで URI add にアクセスする
```html
<form method="post" action="add">
```

- ただし add ボタンを押すとパスが 〜/add になってしまい、このままリロードすると最後の入力が延々とTODOリストに追加されてしまう。

## TodoList に Post-Redirect-Get を追加
- add ボタンを押すとパスが 〜/add になってしまい、このままリロードすると最後の入力が延々とTODOリストに追加されてしまう問題点を、Post-Redirect-Get により解消する
- 〜/add の実行後、〜/todo にリダイレクトさせることで、リロードしても 〜/add が実行されることを防ぐ

```sh
go run ch06/TodoListAddPostRedirectGet/main.go
```
http://localhost:8080/todo

- handleAdd()関数の最後の処理を、handleTodo関数の実行ではなく、todoパスへのRedirectに変える
```sh
func handleAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	todo := r.Form.Get("todo")
	todoList = append(todoList, todo)
	// handleTodo(w, r)
	http.Redirect(w, r, "/todo", 303)
}
```

## セッション管理

go mod init ch06/TodoListSession

```sh
go run ./ch06/TodoListSession/
```
http://localhost:8080/todo

```go
var todoLists = make(map[string][]string)

// make ... map等の初期化
// map ... key と value を持つデータ構造
// map[string] ... key が string
// []string ... value が string の配列
```

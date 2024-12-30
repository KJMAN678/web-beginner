### 初期ファイル作成
```sh
go mod init web-begginer
```

### フォーマッター
```sh
go fmt ./...
```

### ListenAndServe
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

### FileServer static/xxx.html を読み込んでローカルサーバー立ち上げ

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

```sh
fileServer := http.FileServer(http.Dir("ch06/FileServer/static"))
stripPrefix := http.StripPrefix("/static/", fileServer)
http.Handle("/", stripPrefix)
```


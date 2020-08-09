# Go Echo Tutorial
GoのWebフレームワークであるEchoを使うチュートリアル

## modを初期化
```
$ go mod init example.com/go_echo_tutorial
```

## build
buildで依存モジュールを自動インストールする
```
$ go build
```

## server
http://localhost:1323
```
$ go run server.go
```

## module削除
使われていない依存モジュールを削除する
```
$ go mod tidy
```

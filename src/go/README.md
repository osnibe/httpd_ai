# Go

## 実行

```sh
go run main.go
```

## ビルドして実行

```sh
go build -o httpd main.go
./httpd
```

## 停止

`Ctrl+C`

## 動作確認

```sh
curl http://localhost:8080/
# => Hello, World!

# ポートの LISTEN 状態を確認
netstat -an -p tcp | grep 8080

# プロセス名・PID も確認（macOS 推奨）
lsof -i :8080
```

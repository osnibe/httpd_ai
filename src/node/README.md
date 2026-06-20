# Node.js

## 実行

```sh
node main.js
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

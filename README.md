AIを使ってhttpdを作る

## ビルド

```sh
make
```

## 実行

```sh
./httpd
```

ポート 8080 で待ち受けます。

## 動作確認

```sh
curl http://localhost:8080/
# => Hello, World!
```

## クリーンアップ

```sh
make clean
```

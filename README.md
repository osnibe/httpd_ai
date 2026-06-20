# httpd_ai

AIによるコーディング（Claude Code）を活用して、HTTP サーバーをゼロから複数言語で実装するプロジェクト。

## 実装一覧

| 言語 | ディレクトリ |
|------|------------|
| C | [src/c](src/c) |
| Go | [src/go](src/go) |
| Java | [src/java](src/java) |
| Python | [src/python](src/python) |
| Node.js | [src/node](src/node) |

各言語のビルド・実行・停止・動作確認方法は各ディレクトリの README を参照してください。

## 現在の実装状況

- TCP ソケットで接続を受け付け、固定の `HTTP 200 Hello, World!` を返す
- HTTP パース・ルーティングは未実装

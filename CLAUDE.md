# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

このプロジェクトは、AIによるコーディング（Claude Code）を活用して HTTP サーバー (`httpd`) をゼロから作成する実験的プロジェクトです。実装言語・フレームワークは未定。

## Build & Run

```sh
make          # ビルド → ./httpd
make clean    # 生成物を削除
./httpd       # ポート 8080 で起動
```

## Architecture

```
src/main.c        TCPソケット生成・accept ループ・レスポンス送信
Makefile          ビルド定義
```

- `create_server_socket()` — ソケット作成、`SO_REUSEADDR`、bind、listen
- `handle_connection()` — クライアントからリクエストを読み、固定レスポンスを返す（HTTP パース未実装）
- メインループは現状シングルスレッド・逐次処理

## Current Status

TCP 接続の受け付けと固定レスポンス返却のみ実装済み。HTTP パース・ルーティングはこれから。

# httpeek

HTTP レスポンスヘッダを確認するための疎通確認ツールです。

## 概要

`httpeek`は指定された URL に HTTP GET リクエストを送信し、レスポンスのヘッダ情報を表示するコマンドラインツールです。
ボディは読み捨てられ、ヘッダのみが表示されます。

## 機能

- HTTP GET リクエストの送信
- HTTP ステータスコードの表示
- レスポンスヘッダの表示
- レスポンスボディの読み捨て（意図的）
- HTTP リダイレクト非対応（意図的）

## ビルド

### 全プラットフォーム用ビルド（デフォルト）

```bash
make
```

または

```bash
make all
```

### 指定プラットフォームのみのビルド

```bash
make linux
make windows
make darwin
```

## 使用方法

```bash
./httpeek http://www.example.com
```

### 実行例

```bash
$ ./httpeek http://www.example.com
200 OK

Accept-Ranges: bytes
Cache-Control: max-age=604800
Content-Type: text/html; charset=UTF-8
Date: Thu, 21 Aug 2025 11:17:51 GMT
Etag: "3147526947+gzip"
Expires: Thu, 28 Aug 2025 11:17:51 GMT
Last-Modified: Thu, 17 Oct 2019 07:18:26 GMT
Server: ECS (nyb/1D13)
Vary: Accept-Encoding
X-Cache: HIT
Content-Length: 1256
```

## 注意事項

- HTTP リダイレクトには自動的に対応しません
- レスポンスボディは表示されず、読み捨てられます
- HEAD メソッドではなく、GET メソッドを使用します

## ビルド環境

- Go 1.23 以上
- CGO 無効でのスタティックビルド
- クロスプラットフォーム対応（Linux/Windows/macOS）

# httpeek

HTTP レスポンスヘッダとボディを確認するための疎通確認ツールです。

## 概要

`httpeek`は指定された URL に HTTP GET リクエストを送信し、レスポンスのヘッダ情報とボディを表示するコマンドラインツールです。

## 機能

- HTTP GET リクエストの送信
- HTTP ステータスコードの表示
- レスポンスヘッダの表示
- レスポンスボディの表示
- パイプ/リダイレクト時の自動ヘッダー非表示
- ヘッダーのみ表示オプション（-I/--head）
- SSL証明書検証のスキップオプション（-k/--insecure）
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

### 基本的な使用方法

```bash
./httpeek http://www.example.com
```

### オプション

- `-I`, `--head`: ヘッダーのみを表示（ボディは表示しない）
- `-k`, `--insecure`: SSL証明書の検証をスキップ

### 実行例

#### 通常の実行（ヘッダー + ボディを表示）

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

<!doctype html>
<html>
<head>
    <title>Example Domain</title>
...
```

#### ヘッダーのみ表示

```bash
$ ./httpeek -I http://www.example.com
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

#### パイプでJSONデータを処理

パイプやリダイレクトで使用する場合は自動的にボディのみを出力するので、JSONデータの処理などに適しています。

```bash
$ ./httpeek https://api.example.com/data.json | jq .
{
  "name": "example",
  "version": "1.0.0"
}
```

#### ファイルにリダイレクト

```bash
./httpeek https://api.example.com/data.json > data.json
```

## 注意事項

- HTTP リダイレクトには自動的に対応しません
- HEAD メソッドではなく、GET メソッドを使用します
- パイプやリダイレクト使用時は自動的にヘッダーが非表示になります

## ビルド環境

- Go 1.23 以上
- CGO 無効でのスタティックビルド
- クロスプラットフォーム対応（Linux/Windows/macOS）

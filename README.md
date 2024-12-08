# ファイルパスチェッカー

このツールは、指定されたディレクトリ内のファイルパスを検証し、設定されたルールに基づいてチェックを行うGo言語で書かれたプログラムです。

## 機能

- 指定されたディレクトリを再帰的に探索
- 正規表現によるファイル名のマッチング
- パスパターンに基づく検証
- カスタマイズ可能なエラーメッセージ
- Windows/Linux/Mac対応

## 必要条件

- Go 1.x以上
- Git

## セットアップ

```bash
mkdir file-path-checker
cd file-path-checker
go mod init file-path-checker
go get gopkg.in/yaml.v3
go mod tidy
```

## ビルド方法

### Linux/macOS

```bash
go build -o file-checker
```

### Windows

```bash
go build -o file-checker.exe
```

## 実行方法

### Linux/macOS

```bash
./file-checker
```

### Windows

```bash
./file-checker.exe
```

## テスト実行

```bash
go test -v
```

## 設定

`.check_file_path.yaml`ファイルでパスのチェックルールを設定できます。
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

## システムワイドでの実行設定

### Linux/macOS

```bash
# バイナリを/usr/local/binにコピー
sudo cp file-checker /usr/local/bin/

# または、~/.bashrcや~/.zshrcに以下を追加
export PATH=$PATH:/path/to/file-checker
```

### Windows

1. システムのプロパティを開く
2. 環境変数をクリック
3. システム環境変数のPATHを選択して編集
4. file-checkerの実行ファイルがあるディレクトリのパスを追加
5. OKをクリックして保存

設定後は、任意のディレクトリから`file-checker`コマンドを実行できます。

## テスト実行

```bash
go test -v
```

## 設定ファイル

OSに応じて`.check_file_path.linux.yaml`または`.check_file_path.windows.yaml`ファイルでパスのチェックルールを設定できます。以下は設定ファイルの例です：

```yaml
rules:
  - base_path: "./src"
    file_name: "Controller\\.kt$"
    regexes:
      - regex: "^.*/domain/.*$"
        message: "NG: domain folder detected"
      - regex: "^.*/presentation/controller/UserPoolController\\.kt$"
        message: "OK: Valid file path"
  - base_path: "./lib"
    file_name: "Hoge\\.kt$"
    regexes:
      - regex: "^.*/someotherpath/.*$"
        message: "NG: Invalid path"
```

### 設定項目の説明

- `base_path`: チェック対象のベースディレクトリ
  - 検索を開始するルートディレクトリを指定します

- `file_name`: チェック対象のファイル名パターン
  - 正規表現で指定します
  - 例: `Controller\\.kt$`はファイル名が`Controller.kt`で終わるファイルにマッチ

- `regexes`: パスチェックルールのリスト
  - `regex`: パスパターンを正規表現で指定
  - `message`: パターンにマッチした際に表示するメッセージ
    - マッチ結果の説明やエラーメッセージを設定します

このツールは設定ファイルに基づいて、ファイルが適切な場所に配置されているかを正規表現でチェックし、指定されたメッセージを表示します。
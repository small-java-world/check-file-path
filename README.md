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

`.check_file_path.yaml`ファイルでパスのチェックルールを設定できます。以下は設定ファイルの例です：

```yaml
rules:
  - pattern: "src/domain/*.kt"
    description: "ドメインレイヤーのKotlinファイル"
    allowed: true

  - pattern: "src/presentation/controller/*.kt"
    description: "プレゼンテーション層のコントローラー"
    allowed: true

  - pattern: "lib/**/*.kt"
    description: "ライブラリ配下のKotlinファイル"
    allowed: false
```

### 設定項目の説明

- `pattern`: チェック対象のファイルパスパターン（glob形式）
  - `*`: 任意の文字列（ディレクトリ区切り文字を除く）
  - `**`: 任意の階層のディレクトリ

- `description`: ルールの説明文（エラーメッセージに使用）

- `allowed`: 
  - `true`: このパターンに一致するパスを許可
  - `false`: このパターンに一致するパスを禁止

このツールは設定ファイルに基づいて、ファイルパスが適切な場所に配置されているかをチェックします。
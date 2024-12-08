package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"gopkg.in/yaml.v3"
)

// OSに応じた正規表現パターンを返す
func getOSPathPattern(pattern string) string {
	if runtime.GOOS == "windows" {
		return pattern
	}
	// Linuxの場合はバックスラッシュをスラッシュに変換
	return filepath.FromSlash(pattern)
}

func TestMatchFile(t *testing.T) {
	// OSに応じたパターンを設定
	domainPattern := "^.*/domain/.*$"
	userPoolPattern := "^.*/presentation/controller/UserPoolController\\.kt$"

	tests := []struct {
		name     string
		path     string
		rule     Rule
		expected string
	}{
		{
			name: "domainフォルダ内のファイルを検出",
			path: filepath.FromSlash("src/domain/Controller.kt"),
			rule: Rule{
				BasePath: "./src",
				FileName: "Controller\\.kt$",
				Regexes: []struct {
					Regex   string `yaml:"regex"`
					Message string `yaml:"message"`
				}{
					{
						Regex:   domainPattern,
						Message: "NG: domain folder detected",
					},
				},
			},
			expected: filepath.FromSlash("src/domain/Controller.kt") + ": NG: domain folder detected",
		},
		{
			name: "許可されたUserPoolControllerを検出",
			path: filepath.FromSlash("src/presentation/controller/UserPoolController.kt"),
			rule: Rule{
				BasePath: "./src",
				FileName: "Controller\\.kt$",
				Regexes: []struct {
					Regex   string `yaml:"regex"`
					Message string `yaml:"message"`
				}{
					{
						Regex:   userPoolPattern,
						Message: "OK: Valid file path",
					},
				},
			},
			expected: filepath.FromSlash("src/presentation/controller/UserPoolController.kt") + ": OK: Valid file path",
		},
		{
			name: "問題のないパスを検出",
			path: filepath.FromSlash("src/other/Controller.kt"),
			rule: Rule{
				BasePath: "./src",
				FileName: "Controller\\.kt$",
				Regexes: []struct {
					Regex   string `yaml:"regex"`
					Message string `yaml:"message"`
				}{
					{
						Regex:   domainPattern,
						Message: "NG: domain folder detected",
					},
				},
			},
			expected: filepath.FromSlash("src/other/Controller.kt") + ": 問題なし",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchFile(tt.path, tt.rule)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestCheckFiles(t *testing.T) {
	// テスト用の一時ディレクトリを作成
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// テスト用のディレクトリ構造を作成
	dirs := []string{
		filepath.Join(tmpDir, "src", "domain"),
		filepath.Join(tmpDir, "src", "presentation", "controller"),
		filepath.Join(tmpDir, "src", "other"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}

	// テスト用のファイルを作成
	files := []string{
		filepath.Join(tmpDir, "src", "domain", "Controller.kt"),
		filepath.Join(tmpDir, "src", "presentation", "controller", "UserPoolController.kt"),
		filepath.Join(tmpDir, "src", "other", "Controller.kt"),
	}

	for _, file := range files {
		if err := os.WriteFile(file, []byte(""), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// パターンを設定（スラッシュ形式に統一）
	domainPattern := "^.*/domain/.*$"
	userPoolPattern := "^.*/presentation/controller/UserPoolController\\.kt$"

	// テストケース
	rule := Rule{
		BasePath: tmpDir,
		FileName: "Controller\\.kt$",
		Regexes: []struct {
			Regex   string `yaml:"regex"`
			Message string `yaml:"message"`
		}{
			{
				Regex:   domainPattern,
				Message: "NG: domain folder detected",
			},
			{
				Regex:   userPoolPattern,
				Message: "OK: Valid file path",
			},
		},
	}

	// checkFilesの実行
	checkFiles(rule)
}

func TestLoadConfig(t *testing.T) {
	// OSに応じた設定ファイルを選択
	var yamlContent []byte
	if runtime.GOOS == "windows" {
		yamlContent = []byte(`
rules:
  - base_path: "./src"
    file_name: "Controller\\.kt$"
    regexes:
      - regex: "^.*\\\\domain\\\\.*$"
        message: "NG: domain folder detected"
      - regex: "^.*\\\\presentation\\\\controller\\\\UserPoolController\\.kt$"
        message: "OK: Valid file path"
`)
	} else {
		yamlContent = []byte(`
rules:
  - base_path: "./src"
    file_name: "Controller\\.kt$"
    regexes:
      - regex: "^.*/domain/.*$"
        message: "NG: domain folder detected"
      - regex: "^.*/presentation/controller/UserPoolController\\.kt$"
        message: "OK: Valid file path"
`)
	}

	tmpFile := ".check_file_path.yaml.test"
	if err := os.WriteFile(tmpFile, yamlContent, 0644); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile)

	// 設定ファイルを読み込んでパースできることを確認
	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("設定ファイルの読み込みに失敗: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		t.Fatalf("YAMLのパースに失敗: %v", err)
	}

	// 設定内容の検証
	if len(config.Rules) != 1 {
		t.Errorf("期待するルール数: 1, 実際: %d", len(config.Rules))
	}

	expectedBasePath := "./src"
	if config.Rules[0].BasePath != expectedBasePath {
		t.Errorf("期待するBasePath: %s, 実際: %s", expectedBasePath, config.Rules[0].BasePath)
	}
} 
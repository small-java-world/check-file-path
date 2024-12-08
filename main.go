package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v3"
)

// Rule は各チェックルールを定義する構造体です
type Rule struct {
	BasePath string `yaml:"base_path"`
	FileName string `yaml:"file_name"`
	Regexes  []struct {
		Regex   string `yaml:"regex"`
		Message string `yaml:"message"`
	} `yaml:"regexes"`
}

// Config は設定ファイル全体の構造を表します
type Config struct {
	Rules []Rule `yaml:"rules"`
}

func main() {
	// 設定ファイルを読み込む
	data, err := os.ReadFile(".check_file_path.yaml")
	if err != nil {
		log.Fatalf("設定ファイルの読み込みに失敗しました: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("YAMLのパースに失敗しました: %v", err)
	}

	// 各ルールに対してファイルチェックを実行
	for _, rule := range config.Rules {
		fmt.Printf("チェック対象ディレクトリ: %s\n", rule.BasePath)
		checkFiles(rule)
	}
}

func checkFiles(rule Rule) {
	fileNameRegex, err := regexp.Compile(rule.FileName)
	if err != nil {
		log.Printf("不正な正規表現です: %s: %v", rule.FileName, err)
		return
	}

	err = filepath.Walk(rule.BasePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && fileNameRegex.MatchString(info.Name()) {
			result := matchFile(path, rule)
			fmt.Println(result)
		}
		return nil
	})

	if err != nil {
		log.Printf("パスの探索中にエラーが発生しました %s: %v", rule.BasePath, err)
	}
}

func matchFile(path string, rule Rule) string {
	// パスの区切り文字を統一する（スラッシュに変換）
	normalizedPath := filepath.ToSlash(path)
	
	for _, r := range rule.Regexes {
		regex, err := regexp.Compile(r.Regex)
		if err != nil {
			continue
		}
		if regex.MatchString(normalizedPath) {
			return fmt.Sprintf("%s: %s", path, r.Message)
		}
	}
	return fmt.Sprintf("%s: 問題なし", path)
} 
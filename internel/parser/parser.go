package parser

import (
	"log"

	"os"

	"gopkg.in/yaml.v3"
)

// Config 定义对应的结构体
type Config struct {
	DedeUserID      string `yaml:"DedeUserID"`
	DedeUserIDCkMd5 string `yaml:"DedeUserID__ckMd5"`
	SESSDATA        string `yaml:"SESSDATA"`
	BiliJCT         string `yaml:"BiliJCT"`
	Cookie          string `yaml:"Cookie"`
}

func GetConfig(path string) (config Config, err error) {
	// 打开配置文件
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("无法打开配置文件: %v", err)
		return config, err
	}
	defer file.Close()

	// 创建一个解码器
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
		return config, err
	}

	return config, nil
}

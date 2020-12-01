package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	HOST     string `yaml:"host"`
	PORT     string `yaml:"posr"`
	API      string `yaml:"api"`
	TestFile string `yaml:"test_file"`
}

type ImgData struct {
	Name    string
	Content []byte
}

type Global struct {
	Cfg        *Config
	WorkPools  chan string
	StorePools chan *ImgData
	GPUWork    chan bool
}

func NewGlobal() *Global {
	cfg := NewConfig()
	return &Global{
		Cfg:        cfg,
		WorkPools:  make(chan string, 20),
		StorePools: make(chan *ImgData, 20),
		GPUWork   : make(chan bool),
	}
}

func NewConfig() *Config {
	var cfg = &Config{}
	// 读取yaml文件的信息到Config中
	file, err := os.Open("./config/config.yaml")
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("读文件错误: %v", err)
	}
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		fmt.Println("解析yaml失败")
		log.Fatalf("解析config.yaml读取错误: %v", err)
		panic("")
	}
	return cfg
}

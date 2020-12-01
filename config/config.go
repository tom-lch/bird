package config

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct{

}

type ImgData struct {
	Name string
	Content []byte
}

var workPools = make(chan string, 10)
var storePools = make(chan *ImgData, 10)
var concal = make(chan int)

var cfg *Config



func init() {
	NewConfig()
}

func NewConfig() {
	// 读取yaml文件的信息到Config中
	file, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		panic(err)
	}
}

package workPool

import (
	"bird/config"
	"bufio"
	"fmt"
	"io"
	"os"
)

// 从任务池中抓取到图片链接，装入到 workPool 池中
func GetSpiderWork(glb *config.Global) {
	GetFromTxt(glb)
}

// 使用保存在文件中的url放入到 workPool 中，后续调整成从数据库中抓取
func GetFromTxt(glb *config.Global) {
	f, err := os.Open(glb.Cfg.TestFile)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	totLine := 0
	for {
		content, isPrefix, err := reader.ReadLine()
		fmt.Println(string(content), isPrefix, err)
		glb.WorkPools <- string(content)
		if !isPrefix {
			totLine++
		}
		if err == io.EOF {
			fmt.Println("一共有", totLine, "行内容")
			close(glb.WorkPools)
			break
		}
	}
}

func GetFromGSE() error {
	// DB := connectDB()
	// select URL from DBName where 抓取规则后续调整
	return nil
}

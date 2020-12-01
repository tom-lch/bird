package spider

import (
	"bird/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// config.cfg
// 爬虫的任务就是根据从任务池拿到的任务然后进行下载图片

// 目前有两个通道，一个是 workPool 传递下载任务， 一个是 storePool 保存下载图片的路径

func GetWorkFromChan(glb *config.Global) {
	for url := range glb.WorkPools {
		DLPhoto(url, glb)
	}
	close(glb.StorePools)
}

func CreateName() string {
	return fmt.Sprintf("img/v%d.jpg", time.Now().Unix())
}

func DLPhoto(url string, glb *config.Global) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// 在此处调用GPU去处理
	photoname := CreateName()
	info := &config.ImgData{Name: photoname, Content: body}
	glb.StorePools <- info
	return ioutil.WriteFile(photoname, body, 0755)
}

package useGPU

import (
	"bird/config"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

// 调用GPU处理图片
// 调用方式1 使用http发起本地post请求调用 localhost 适用于本地部署
// 调用方式2 使用http 发起局域网内部调用 192.168.2.229:8868 调用
func GetOCRInfo(glb *config.Global) {
	// 从 storePools 获取到img的[]byte 格式
	// httpPostForm(ImgBase64)
	var wg sync.WaitGroup
	for info := range glb.StorePools {
		wg.Add(1)
		// go ConnectGPUByFile(info.Name, glb)
		go func(info *Config.ImgData, glb *config.Global) {
			httpPostForm(Byte2Base64(info.Content), glb)
			wg.Wait()
		}(info, glb)
	}
}

func ConnectGPUByFile(imgFile string, glb *config.Global) {
	imageBase64 := base64ImgByfile(imgFile)
	httpPostForm(imageBase64, glb)
}

func Byte2Base64(code []byte) string {
	return base64.StdEncoding.EncodeToString(code)
}

func base64ImgByfile(imgFile string) string {
	image, _ := ioutil.ReadFile(imgFile)
	imageBase64 := base64.StdEncoding.EncodeToString(image)
	return imageBase64
}

// {"msg":"","results":[[{"confidence":0.8403433561325073,"text":"约定","text_region":[[345,377],[641,390],[634,540],[339,528]]},{"confidence":0.8131805658340454,"text":"最终相遇","text_region":[[356,532],[624,530],[624,596],[356,598]]}]],"status":"0"}

func httpPostForm(ImgBase64 string, glb *config.Global) (string, error) {
	resp, err := http.PostForm(glb.Cfg.HOST+glb.Cfg.PORT+glb.Cfg.API, url.Values{"images": {ImgBase64}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))
	return string(body), nil
}

package useGPU

import (
	"fmt"
	"net/http"
	"bird/config"
	"encoding/base64"
	"net/url"
)

// 调用GPU处理图片
// 调用方式1 使用http发起本地post请求调用 localhost 适用于本地部署
// 调用方式2 使用http 发起局域网内部调用 192.168.2.229:8868 调用
func GetOCRInfo() {
	// 从 storePools 获取到img的[]byte 格式
	// httpPostForm(ImgBase64)
	for info := range config.storePools {
		// go ConnectGPUByFile(info.Name)
		go httpPostForm(Byte2Base64(info.Content))
	}
}

func ConnectGPUByFile(imgFile) {
	imageBase64 := base64ImgByfile(imgFile)
	httpPostForm(ImgBase64)
}

func Byte2Base64(code) string {
	return base64.StdEncoding.EncodeToString(code)
}

func base64ImgByfile(imgFile) string {
	image, _:= ioutil.ReadFile(imgFile)
	imageBase64 := base64.StdEncoding.EncodeToString(image)
	return imageBase64
}

// {"msg":"","results":[[{"confidence":0.8403433561325073,"text":"约定","text_region":[[345,377],[641,390],[634,540],[339,528]]},{"confidence":0.8131805658340454,"text":"最终相遇","text_region":[[356,532],[624,530],[624,596],[356,598]]}]],"status":"0"}

func httpPostForm(ImgBase64) (string error) {
	resp, err := http.PostForm(config.cfg.Host+config.cfg.Port+config.cfg.API, url.Values{"images": {ImgBase64}})
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




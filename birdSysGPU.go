package main

import (
	"time"

	"bird/pkg/spider"
	"bird/config"
	"bird/pkg/useGPU"
	"bird/pkg/workPool"

	"sync/WaitGroup"
)

// 实现多线程从任务池拿去任务，多线程爬虫下载图片，多线程任务调度获取GPU资源，保障GPU满负荷运行
func main() {
	new :=time.Now()
	fmt.Println(new.Format("2006-01-02 15:04:05"))
	
	start := time.Now().UnixNano()
	defer func(){
		// 关闭通道
		config.storePools.Close()
	}()
	var wg WaitGroup
	// // 从任务池获取任务
	// go workPool.GetSpiderWork()
	// // 使用爬虫下载图片
	// go spider.GetWorkFromGSE()
	// // 使用gpu设备处理
	// go useGPU.GetOCRInfo()
	opts := []func(){workPool.GetSpiderWork, spider.GetWorkFromChan, useGPU.GetOCRInfo}
	for _, opt := range opts {
		wg.Add(1)
		go func(){
			opt()
		}()
		wg.Done()
	}
	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Printf("cost is :%d \n", ( end-start)/1000)
}
package main

import (
	"fmt"
	"kkAndroidPackServer/config"
	"kkAndroidPackServer/db/bean"
	"kkAndroidPackServer/db/sql"
	httpServer "kkAndroidPackServer/http"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	mysql.InitMySQL()
	apps := bean.FetchUnPackTask()
	fmt.Println(apps)

	var stopLock sync.Mutex
	stop := false
	stopChan := make(chan struct{}, 1)
	signalChan := make(chan os.Signal, 1)
	go func() {
		//阻塞程序运行，直到收到终止的信号
		<-signalChan
		stopLock.Lock()
		stop = true
		stopLock.Unlock()
		fmt.Println("程序要退出了")
		log.Println("Cleaning before stop...")
		stopChan <- struct{}{}
		os.Exit(0)
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	handler := httpServer.HandleHTTPServer("/")
	url := config.HttpHost + ":" + config.HttpPort
	httpServer.Start(url, handler)

}

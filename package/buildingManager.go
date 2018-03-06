package pack

import (
	"kkAndroidPackServer/db/bean"
	"sync"
	"time"
)

type Manager struct {
	lastBuildNumber int
	status          string
}

var instance *Manager
var once sync.Once

//Instance 获取对象
func Instance() *Manager {
	once.Do(func() {
		instance = &Manager{}

		go startTimer()
	})
	return instance
}

func startTimer() {
	checkTimeOutTask()

	timer := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-timer.C:
			checkTimeOutTask()
		}
	}
}

func checkTimeOutTask() {
	apps := bean.FetchTimeOutBuildingPackTask()
	if len(apps) > 0 {
		for _, x := range apps {
			app := x.(*bean.PackageApp)
			app.Status = "waiting"
			app.StartTime = time.Now().Format("2006-01-02 15:04:05")
			app.FinishTime = time.Now().Format("2006-01-02 15:04:05")
			app.Update()
		}
	}
}

package pack

import (
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
	timer := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-timer.C:
			dealPackage()
		}
	}

	dealPackage()
}

func dealPackage() {
	for {

	}
}

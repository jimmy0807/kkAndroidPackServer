package bean

import (
	"kkAndroidPackServer/db/sql"
	"log"
	"os"
	"time"
)

type PackageApp struct {
	mysql.MySQLReflect

	ApkName           string `json:"apk_name"`
	ApkLastUpdateTime string `json:"apk_last_update_time"`
	Status            string `json:"status"`
	ChannelName       string `json:"channel_name"`
	ChannelID         int64  `json:"channel_id"`
}

//GenerateReflectValue 实例
func (p *PackageApp) GenerateReflectValue() interface{} {
	return &PackageApp{}
}

//Insert 插入
func (p *PackageApp) Insert() {
}

func GetFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

//Next 遍历
func (p *PackageApp) Next(n interface{}) {
	app := n.(*PackageApp)
	file, err := os.Open("./" + app.ApkName)
	if err != nil {
		return
	}

	fi, err := file.Stat()
	if err != nil {
		return
	}

	app.ApkLastUpdateTime = fi.ModTime().Format("2006-01-02 15:04:05")
}

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
	WriteTime         string `json:"write_time"`
	HostName          string `json:"host_name"`
	StartTime         string `json:"start_time"`
	FinishTime        string `json:"finish_time"`
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

func (p *PackageApp) Update() {
	p.WriteTime = time.Now().Format("2006-01-02 15:04:05")
	mysql.Exec("Update channel Set status = ?, write_time = STR_TO_DATE(?, '%Y-%m-%d %H:%i:%s'),host_name = ?, start_time = STR_TO_DATE(?, '%Y-%m-%d %H:%i:%s'), finish_time = STR_TO_DATE(?, '%Y-%m-%d %H:%i:%s') Where channel_id = ?", p.Status, p.WriteTime, p.HostName, p.StartTime, p.FinishTime, p.ChannelID)
}

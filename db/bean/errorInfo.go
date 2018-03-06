package bean

import (
	"kkAndroidPackServer/db/sql"
	"time"
)

type ErrorInfo struct {
	mysql.MySQLReflect

	Description string `json:"description"`
	ID          string `json:"id"`
	HostName    string `json:"host_name"`
}

//GenerateReflectValue 实例
func (p *ErrorInfo) GenerateReflectValue() interface{} {
	return &ErrorInfo{}
}

//Insert 插入
func (p *ErrorInfo) Insert() {
	mysql.Insert("INSERT INTO error(description,host_name,write_time) VALUES(?,?,STR_TO_DATE(?, '%Y-%m-%d %H:%i:%s'),STR_TO_DATE(?, '%Y-%m-%d %H:%i:%s'),?,?)", p.Description, p.HostName, time.Now().Format("2006-01-02 15:04:05"))
}

func (p *ErrorInfo) Update() {
	//p.WriteTime = time.Now().Format("2006-01-02 15:04:05")
	//mysql.Exec("Update channel Set status = ?, write_time = STR_TO_DATE(?, '%Y-%m-%d %H:%i:%s'),host_name = ? Where channel_id = ?", p.Status, p.WriteTime, p.HostName, p.ChannelID)
}

package bean

import (
	"kkAndroidPackServer/db/sql"
)

type PackageDir struct {
	mysql.MySQLReflect

	Dir string `json:"dir"`
}

//GenerateReflectValue 实例
func (p *PackageDir) GenerateReflectValue() interface{} {
	return &PackageDir{}
}

//Insert 插入
func (p *PackageDir) Insert() {
}

func (p *PackageDir) Next(n interface{}) {
}

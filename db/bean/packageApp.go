package bean

import (
	"kkAndroidPackClient/db/sql"
)

type PackageApp struct {
	mysql.MySQLReflect

	Name           string      `json:"name"`
	DisplayName    string      `json:"display_name"`
	ID             int64       `json:"id"`
	SortIndex      int64       `json:"sort_index"`
	DictionaryName string      `json:"-"`
	Os             string      `json:"-"`
	Branch         interface{} `json:"branch"`
	JenkinsURL     string      `json:"-"`
}

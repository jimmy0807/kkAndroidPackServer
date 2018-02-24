package handler

import (
	"kkAndroidPackServer/db/bean"
	"net/http"
)

//HandlerGetPackageList 处理打包请求
func HandlerGetPackageList(w http.ResponseWriter, r *http.Request) {
	apps := bean.FetchUnPackTask()
	writeJSONResponse(w, apps)
}

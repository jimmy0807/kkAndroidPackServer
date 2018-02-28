package handler

import (
	"kkAndroidPackServer/db/bean"
	"net/http"
)

//HandlerGetPackageList 处理打包请求
func HandlerGetPackageList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	params := r.Form
	hostName := params["hostName"][0]

	apps := bean.FetchUnPackTask()
	if len(apps) == 1 {
		app := apps[0].(*bean.PackageApp)
		app.Status = "building"
		app.HostName = hostName
		app.Update()
	}
	writeJSONResponse(w, apps)
}

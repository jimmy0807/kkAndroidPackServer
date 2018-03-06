package handler

import (
	"fmt"
	"kkAndroidPackServer/db/bean"
	"net/http"
	"time"
)

//HandlerGetPackageList 处理打包请求
func HandlerGetPackageList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandlerGetPackageList")
	r.ParseForm()
	params := r.Form
	hostName := params["hostName"][0]

	apps := bean.FetchUnPackTask()
	if len(apps) == 1 {
		app := apps[0].(*bean.PackageApp)
		app.Status = "building"
		app.HostName = hostName
		app.StartTime = time.Now().Format("2006-01-02 15:04:05")
		app.FinishTime = time.Now().Format("2006-01-02 15:04:05")
		app.Update()
	}
	writeJSONResponse(w, apps)
	fmt.Println("HandlerGetPackageList writeJSONResponse")
}

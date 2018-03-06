package handler

import (
	"kkAndroidPackServer/db/bean"
	"net/http"
)

//HandlerGetPackageList 处理打包请求
func HandlerError(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	params := r.Form
	hostName := params["hostName"][0]
	desc := params["desc"][0]

	e := new(bean.ErrorInfo)
	e.HostName = hostName
	e.Description = desc

	writeJSONResponse(w, new(interface{}))
}

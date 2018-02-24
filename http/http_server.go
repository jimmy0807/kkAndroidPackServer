package http

import (
	"kkAndroidPackServer/http/handler"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const uploadPath = "./"

//HandleHTTPServer 水电费
func HandleHTTPServer(rootPath string) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/uploadApkFile", handler.UploadApkFileHandler).Methods("POST")
	r.HandleFunc("/fetchPackTask", handler.HandlerGetPackageList).Methods("GET")

	fs := http.FileServer(http.Dir(uploadPath))
	//如果是http.hander的话 不用加{rest}
	r.Handle("/files/{rest}", http.StripPrefix("/files", fs))

	//r.HandleFunc("/api/doPackage", handler.HandlerDoPackage).Methods("GET")
	//r.HandleFunc("/api/test", handler.HandlerGetPackageList).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 	zipName := "test.zip"
		// 	file, _ := os.Open("./test.zip")
		// 	// fileName := path.Base(fileFullPath)
		// 	// fileName = url.QueryEscape(fileName)
		// 	w.Header().Set("Content-Type", "application/octet-stream")
		// 	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipName))
		// 	io.Copy(w, file)
		w.Write([]byte("Wait till you sddee me in action!"))
	})

	http.Handle(rootPath, r)
	return r
}

//Start 初始化
func Start(url string, handler http.Handler) {
	headsOpts := handlers.AllowedHeaders([]string{})
	http.ListenAndServe(url, handlers.CORS(headsOpts)(handler))
}

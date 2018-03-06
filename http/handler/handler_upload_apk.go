package handler

import (
	"fmt"
	"io/ioutil"
	"kkAndroidPackServer/db/bean"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var uploadPath = "./"

func UploadApkFileHandler(w http.ResponseWriter, r *http.Request) {
	//设置文件大小限制
	/*
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}
	*/
	channelID, err := strconv.ParseInt(r.PostFormValue("channelID"), 10, 64)
	fileName := r.PostFormValue("fileName")
	hostName := r.PostFormValue("hostName")
	fmt.Println(hostName)

	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		//enderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		//renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	if filetype != "image/jpeg" && filetype != "image/jpg" &&
		filetype != "image/gif" && filetype != "image/png" &&
		filetype != "application/zip" {
		//renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return
	}

	// fileEndings, err := mime.ExtensionsByType(filetype)
	// if err != nil {
	// 	//renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
	// 	return
	// }

	//newPath := filepath.Join(uploadPath, fileName+fileEndings[0])

	dirs := bean.FetchPackDir()
	if len(dirs) == 1 {
		dir := dirs[0].(*bean.PackageDir)
		uploadPath = dir.Dir
	}

	newPath := filepath.Join(uploadPath, fileName)
	fmt.Printf("FileType: %s, File: %s\n", filetype, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		//renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()
	if _, err := newFile.Write(fileBytes); err != nil {
		//renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("SUCCESS"))
	updateTaskStatus(channelID, "success", hostName)
}

func updateTaskStatus(channelID int64, status string, hostName string) {
	tasks := bean.FetchPackageTaskbyID(channelID)
	if len(tasks) == 1 {
		task := tasks[0].(*bean.PackageApp)
		task.Status = status
		task.HostName = hostName
		task.FinishTime = time.Now().Format("2006-01-02 15:04:05")
		task.Update()
	}
}

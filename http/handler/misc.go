package handler

import (
	"encoding/json"
	"net/http"
)

// JSONResponse 返回数据
type JSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	jsonResponse := &JSONResponse{
		Code:    0,
		Message: "",
		Data:    data,
	}

	bytesJSON, _ := json.Marshal(jsonResponse)

	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(bytesJSON)
}

func writeErrorJSONResponse(w http.ResponseWriter, statusCode int, code int, message string) {
	jsonResponse := &JSONResponse{
		Code:    code,
		Message: message,
	}

	bytesJSON, _ := json.Marshal(jsonResponse)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	w.Write(bytesJSON)
}

func DecodeJSONResponse(body []byte, data interface{}) *JSONResponse {
	jsonResponse := &JSONResponse{
		Code:    0,
		Message: "",
		Data:    data,
	}

	if err := json.Unmarshal(body, &jsonResponse); err == nil {
		// fmt.Println(dat)
		// fmt.Println(dat["result"])
	} else {
		//fmt.Println(err)
	}

	return jsonResponse
}

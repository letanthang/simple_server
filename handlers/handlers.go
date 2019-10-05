package handlers

import (
	"net/http"
)

func CheckHealth(writer http.ResponseWriter, req *http.Request) {
	msg := "OK"
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(msg))
}

func GetStudentByID() {

}

func GetStudent() {

}

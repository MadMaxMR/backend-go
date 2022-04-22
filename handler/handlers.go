package handler

import (
	"encoding/json"
	"github.com/MadMaxMR/backend-go/models"
	"log"
	"net/http"
	"strconv"
)

func SendSuccess(w http.ResponseWriter, req *http.Request, status int, model interface{}) {

	json, _ := json.Marshal(model)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(json)))
	content := w.Header().Get("Content-Length")
	w.WriteHeader(status)
	w.Write(json)
	log.Print("'", req.Method, " - ", req.URL.Path, " - ", req.Proto, "' - ", status, " - ", content)
}

func SendSuccessMessage(w http.ResponseWriter, req *http.Request, status int, message string) {
	var data models.Data = models.Data{Message: make(map[string]string)}
	data.Message["success"] = message
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(json)))
	content := w.Header().Get("Content-Length")
	w.WriteHeader(status)
	w.Write(json)
	log.Print("'", req.Method, " - ", req.URL.Path, " - ", req.Proto, "' - ", status, " - ", content)
}

func SendFail(w http.ResponseWriter, req *http.Request, status int, err string) {
	var data models.Data = models.Data{Message: make(map[string]string)}
	data.Message["error"] = err
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(json)))
	content := w.Header().Get("Content-Length")
	w.WriteHeader(status)
	w.Write(json)
	log.Print("'", req.Method, " - ", req.URL.Path, " - ", req.Proto, " - ", status, " - ", content)
}

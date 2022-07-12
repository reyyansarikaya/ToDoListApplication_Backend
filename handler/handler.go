package handler

import (
	"bootcamp/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IHandler interface {
	HandlerEndpoints(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IService
}

func NewHandler(service service.IService) IHandler {
	return &Handler{service: service}
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	httpErr := HTTPError{
		Code:    code,
		Message: msg,
	}
	bytesResponse, _ := json.Marshal(httpErr)
	w.WriteHeader(code)
	w.Write(bytesResponse)
}
func (h *Handler) HandlerEndpoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == http.MethodGet {
		h.Get(w)
	} else if r.Method == http.MethodPost {
		h.Save(w, r)
	}
}

func (h *Handler) Get(w http.ResponseWriter) {
	todolist := h.service.Get()

	response, err := json.Marshal(todolist)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	c := make(map[string]string)
	err = json.Unmarshal(body, &c)

	todo, ok := c["todo"]
	result := h.service.Save(todo)

	if ok != true || result != 200 || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

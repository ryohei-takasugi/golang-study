package handler

import "net/http"

type EchoHandler struct {
	response http.ResponseWriter
	request  *http.Request
}

func CreateEchoHandler(w http.ResponseWriter, r *http.Request) *EchoHandler {
	return &EchoHandler{
		response: w,
		request:  r,
	}
}

func (h EchoHandler) Handle() {
	h.response.WriteHeader(http.StatusOK)
	h.response.Write([]byte("Hello, Echo! Path:" + h.request.URL.Path))
}

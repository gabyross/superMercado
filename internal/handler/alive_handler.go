package handler

import "net/http"

type AliveHandler struct {
}

func NewAliveHandler() *AliveHandler {
	return &AliveHandler{}
}
func (ah *AliveHandler) Alive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

package controllers

import "net/http"

const (
	pong = "pong"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Pong(w http.ResponseWriter, r *http.Request)
}

type pingController struct{}

func (c *pingController) Pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}

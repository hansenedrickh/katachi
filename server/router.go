package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hansenedrickh/katachi/handler"
	"github.com/hansenedrickh/katachi/helper"
)

func Router(router *mux.Router, h handler.Handler) {
	helper.AddRoute(router, http.MethodGet, "/ping", h.PingHandler)
}
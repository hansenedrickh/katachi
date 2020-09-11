package helper

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AddRoute(router *mux.Router, method, pattern string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	if len(middlewares) == 0 {
		router.HandleFunc(pattern, handler).Methods(method)
		return
	}

	var chainedMiddleWares http.Handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		if i == len(middlewares) - 1 {
			chainedMiddleWares = middlewares[i](handler)
		} else {
			chainedMiddleWares = middlewares[i](chainedMiddleWares)
		}
	}

	router.HandleFunc(pattern, chainedMiddleWares.ServeHTTP).Methods(method)
}

package server

import (
	"net/http"

	"github.com/hansenedrickh/katachi/config"
	"github.com/hansenedrickh/katachi/dependencies"
	"github.com/hansenedrickh/katachi/server/api"
	"github.com/hansenedrickh/katachi/server/handler"
	"github.com/hansenedrickh/katachi/server/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(router *mux.Router, cfg config.Config, deps *dependencies.Dependencies) {
	router.Use(middleware.Context)

	pingHandler := api.NewPingHandler(deps.PingUsecase)
	router.Methods(http.MethodGet).Path("/v1/ping").Handler(pingHandler)

	loginFormHandler := handler.NewLoginFormHandler()
	router.Methods(http.MethodGet).Path("/login").Handler(loginFormHandler)

	loginHandler := handler.NewLoginHandler(deps.UserUsecase)
	router.Methods(http.MethodPost).Path("/login").Handler(loginHandler)

	logoutHandler := handler.NewLogoutHandler()
	router.Methods(http.MethodGet).Path("/logout").Handler(logoutHandler)

	registerFormHandler := handler.NewRegisterFormHandler()
	router.Methods(http.MethodGet).Path("/register").Handler(registerFormHandler)

	registerHandler := handler.NewRegisterHandler(deps.UserUsecase)
	router.Methods(http.MethodPost).Path("/register").Handler(registerHandler)

	sampleRouter := router.NewRoute().Subrouter()
	authMiddleware := middleware.Auth(cfg.JWTSecretToken)
	sampleRouter.Use(authMiddleware)

	sampleHandler := handler.NewSampleIndexHandler(deps.SampleUsecase)
	sampleRouter.Methods(http.MethodGet).Path("/samples").Handler(sampleHandler)

	sampleFormHandler := handler.NewSampleFormHandler(deps.SampleUsecase)
	sampleRouter.Methods(http.MethodGet).Path("/samples/form/{id}").Handler(sampleFormHandler)

	sampleInsertHandler := handler.NewSampleInsertHandler(deps.SampleUsecase)
	sampleRouter.Methods(http.MethodPost).Path("/samples").Handler(sampleInsertHandler)

	sampleUpdateHandler := handler.NewSampleUpdateHandler(deps.SampleUsecase)
	sampleRouter.Methods(http.MethodPatch).Path("/samples/{id}").Handler(sampleUpdateHandler)

	sampleDeleteHandler := handler.NewSampleDeleteHandler(deps.SampleUsecase)
	sampleRouter.Methods(http.MethodDelete).Path("/samples/{id}").Handler(sampleDeleteHandler)
}

package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"github.com/hansenedrickh/katachi/common"
	"github.com/hansenedrickh/katachi/config"
	"github.com/hansenedrickh/katachi/handler"
	"github.com/hansenedrickh/katachi/repository"
	"github.com/hansenedrickh/katachi/usecase"
)

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != nil {
		common.Log.Fatal(err.Error())
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	common.Log.Info("Tenkai server shutting down")
	// Finish all apis being served and shutdown gracefully
	apiServer.Shutdown(context.Background())
	common.Log.Info("Tenkai server shutdown complete")
}

func StartAPIServer() {
	common.Log.Info("Starting Tenkai")

	db := config.InitDB()
	defer db.Close()

	r := repository.NewRepository(db)
	u := usecase.NewUsecase(r)
	h := handler.NewHandler(u)

	router := mux.NewRouter()
	Router(router, h)

	handlerFunc := router.ServeHTTP
	n := negroni.New(negroni.NewRecovery())
	n.Use(httpStatLogger())
	n.UseHandlerFunc(handlerFunc)

	portInfo := ":" + strconv.Itoa(config.Port())
	server := &http.Server{Addr: portInfo, Handler: n}
	go listenServer(server)
	common.Log.Info("Started Tenkai")
	waitForShutdown(server)
}

func httpStatLogger() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now()
		next(rw, r)
		responseTime := time.Now()
		deltaTime := responseTime.Sub(startTime).Seconds() * 1000

		if r.URL.Path != "/ping" {
			common.Log.WithFields(logrus.Fields{
				"RequestTime":   startTime.Format(time.RFC3339),
				"ResponseTime":  responseTime.Format(time.RFC3339),
				"DeltaTime":     deltaTime,
				"RequestUrl":    r.URL.Path,
				"RequestMethod": r.Method,
				"RequestProxy":  r.RemoteAddr,
				"RequestSource": r.Header.Get("X-FORWARDED-FOR"),
			}).Debug("HTTP Logs")
		}
	})
}

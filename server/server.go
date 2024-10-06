package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hansenedrickh/katachi/config"
	"github.com/hansenedrickh/katachi/dependencies"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func Start(cfg config.Config) error {
	deps := &dependencies.Dependencies{}
	err := dependencies.Setup(deps, cfg)
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	SetupRouter(router, cfg, deps)

	address := fmt.Sprintf(":%d", cfg.Port)
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	logrus.WithField("port", cfg.Port).Info("Start API server")

	go startServer(server)
	waitForTermination()
	stopServer(server)

	return nil
}

func startServer(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.WithError(err).Fatal("could not start server")
	}
}

func stopServer(server *http.Server) {
	logrus.Infof("waiting for 1 until no more traffic coming. Press Ctrl+C to skip waiting.")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-time.After(1 * time.Second):
	case <-sig:
		logrus.Warning("forcing shutdown")
	}
	signal.Stop(sig)

	logrus.Info("shutting down application server")
	_ = server.Shutdown(context.Background())
}

func waitForTermination() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

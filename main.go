package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yincongcyincong/BasicStudy/bootstrap"
	"github.com/yincongcyincong/BasicStudy/httpserver"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// parse config
	bootstrap.InitConf()

	// init database
	bootstrap.InitDB()

	// init log
	bootstrap.InitLog()

	// register routers
	router := httpserver.InitRouter()

	srv := &http.Server{
		Addr:    bootstrap.AppConfigInstance.HTTPListen,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("Server forced to shutdown:", err)
	}
	logrus.Warningf("Server exiting")
}

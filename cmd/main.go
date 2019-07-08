package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/nandaryanizar/golang-webservice-example/config/registry"
	"github.com/nandaryanizar/golang-webservice-example/config/routing"
	"github.com/nandaryanizar/golang-webservice-example/internal/app/logging"
	"github.com/nandaryanizar/golang-webservice-example/internal/app/middlewares"
	"github.com/sarulabs/di"
)

func main() {
	defer logging.Logger.Sync()

	err := godotenv.Load("config/config.yaml")
	if err != nil {
		logging.Logger.Fatal("Unable to find config.yaml file")
	}

	app, err := registry.NewContainer()
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
	defer app.Ctn.Delete()

	baseMiddleware := func(h http.Handler) http.Handler {
		return middlewares.PanicRecoveryMiddleware(
			di.HTTPMiddleware(h.ServeHTTP, app.Ctn, func(msg string) {
				logging.Logger.Error(msg)
			}),
			logging.Logger,
		)
	}

	r := routing.NewRouter()
	r.Use(baseMiddleware)
	r.Use(middlewares.JwtAuthentication)

	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "80"
	}

	wt, _ := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	if wt == 0 {
		wt = 15
	}

	rt, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	if rt == 0 {
		rt = 15
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         host + ":" + port,
		WriteTimeout: time.Duration(wt) * time.Second,
		ReadTimeout:  time.Duration(rt) * time.Second,
	}

	logging.Logger.Info("Listening on port " + os.Getenv("SERVER_PORT"))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Logger.Error(err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	logging.Logger.Info("Stopping the http server")

	if err := srv.Shutdown(ctx); err != nil {
		logging.Logger.Error(err.Error())
	}
}

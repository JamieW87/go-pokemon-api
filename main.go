package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go-pokemon-api/config"
	"go-pokemon-api/handlers"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {

	env, err := config.Get()
	if err != nil {
		panic(err)
	}

	log := logrus.New()
	muxRouter := mux.NewRouter()

	handlers.Router(muxRouter, log, env)

	h := &http.Server{
		Addr:    ":" + strconv.Itoa(env.Port),
		Handler: muxRouter,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	log.Info(fmt.Sprintf("🌎 Server running on :%d\n", env.Port))

	// run server in new go routine to allow app shutdown signal wait below
	go func() {
		err := h.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error("An error occurred", err)
			log.Error("Terminating program", nil)
			os.Exit(1)
		}
	}()

	// wait for app shutdown message before attempting to close server gracefully
	<-stop

	log.Info("Shutting down server...")

	timeout := time.Duration(5) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = h.Shutdown(ctx)
	if err != nil {
		log.Error("Failed to shutdown gracefully", err)
	} else {
		log.Info("Server shutdown gracefully")
	}

}

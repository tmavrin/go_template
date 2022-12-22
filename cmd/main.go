package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/teltech/logger"
	internalHTTP "github.com/tmavrin/go_template/internal/http"
)

func main() {

	ctx := context.Background()
	log := logger.New()

	log.Info("starting http server")

	server, err := internalHTTP.NewServer(ctx, log)
	if err != nil {
		log.Fatalf("initialize API http server: %s", err)
	}

	log.Info("http service successfully started")

	defer server.Close()

	errChannel := make(chan error, 1)
	go func() {
		errChannel <- server.Start()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-c:
		log.Infof("signal received: %+v, closing gracefully http server", sig)

		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			log.Fatalf("failed to shut down the http server: %s", err)
		}

		log.Infof("the http server was shutdown gracefully")
	case err := <-errChannel:
		if err != nil {
			log.Fatalf("failed to start the http server: %s", err)
		}
	}

}

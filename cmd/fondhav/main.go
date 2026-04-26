package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/N30A/fondhav/internal/config"
	"github.com/N30A/fondhav/internal/db"
	"github.com/N30A/fondhav/internal/handlers"
	"github.com/N30A/fondhav/internal/repository"
	"github.com/N30A/fondhav/internal/routes"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	pool, err := db.Connect(ctx, cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	fundRepo := repository.NewFundRepository(pool)
	fundHandler := handlers.NewFundHandler(fundRepo)

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux, fundHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("listening on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-signalChan
	log.Println("received interrupt, graceful shutdown...")

	serverCtx, serverCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer serverCancel()
	server.Shutdown(serverCtx)

	pool.Close()
	cancel()
}

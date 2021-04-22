package main

import (
	"context"
	"gRPC_test/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	srv := handlers.New()
	srv.Start()

	<-stop
	log.Println("Server shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.Stop(ctx)
}

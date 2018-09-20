package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type shutdownArgs struct {
	Timeout time.Duration
	Server  *http.Server
}

func gracefulShutdown(args shutdownArgs) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), args.Timeout)
	defer cancel()

	if err := args.Server.Shutdown(ctx); err != nil {
		//		logger.Printf("Error: %v\n", err)
		fmt.Printf("Error: %v\n", err)
	} else {
		//		logger.Println("Server stopped")
		fmt.Println("Server stopped")
	}
}

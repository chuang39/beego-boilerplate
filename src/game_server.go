package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Main context with stopping capabilities
	mainCtx, mainCtxCancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	web.Run()

	// Blocking main and waiting for shutdown.
	select {
	case <-mainCtx.Done():
		logger.Info("main : Terminate signal catched")
		logger.Info("main : Start shutdown...")

		//Call stop function to unregisters the signal behavior
		mainCtxCancel()

		// Create context for Shutdown call.
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// placeholder for server shutdown
	}
}

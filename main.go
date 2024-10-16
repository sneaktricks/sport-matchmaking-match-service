package main

import (
	"context"
	"flag"
	"fmt"
	stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sneaktricks/sport-matchmaking-match-service/dal"
	"github.com/sneaktricks/sport-matchmaking-match-service/database"
	"github.com/sneaktricks/sport-matchmaking-match-service/handler"
	"github.com/sneaktricks/sport-matchmaking-match-service/log"
	"github.com/sneaktricks/sport-matchmaking-match-service/router"
	"github.com/sneaktricks/sport-matchmaking-match-service/store"
)

var port = flag.Int("port", 8080, "Server port")

func main() {
	// Parse port flag
	flag.Parse()

	// Connect to DB
	if _, err := database.Initialize(); err != nil {
		stdlog.Fatalf("Failed to initialize database: %s", err.Error())
	}

	// Create stores
	matchStore := store.NewGormMatchStore(dal.Q)

	// Create router and handler
	r := router.New()
	g := r.Group("")
	h := handler.New(matchStore)

	// Register routes to router main group
	h.RegisterRoutes(g)

	// Prepare for graceful shutdown
	go listenForShutdownSignal(func() {
		r.Shutdown(context.Background())
	})

	// Start the server
	err := r.Start(fmt.Sprintf(":%d", *port))
	if err != http.ErrServerClosed {
		r.Logger.Fatal(err)
	}
}

// Listens for an `os.Interrupt` or `SIGTERM` signal
// and runs the provided `shutdownAction` when received.
func listenForShutdownSignal(shutdownAction func()) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	<-s

	log.Logger.Info("Shutdown signal received, shutting down...")
	shutdownAction()
}

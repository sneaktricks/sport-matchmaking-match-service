package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/sneaktricks/sport-matchmaking-match-service/handler"
	"github.com/sneaktricks/sport-matchmaking-match-service/router"
)

var port = flag.Int("port", 8080, "Server port")

func main() {
	flag.Parse()

	r := router.New()
	g := r.Group("")
	h := handler.New()

	h.RegisterRoutes(g)

	err := r.Start(fmt.Sprintf(":%d", *port))
	if err != http.ErrServerClosed {
		r.Logger.Fatal(err)
	}
}

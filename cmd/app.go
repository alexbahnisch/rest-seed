package main

import (
	"app/internal/handlers"
	"app/internal/server"
	"flag"
	"net/http"
)

func main() {
	port := flag.Int("port", 9000, "Application network port.")

	app := server.New()

	app.RegisterHandler("/version", []string{http.MethodGet}, handlers.HandleVersion)
	app.RegisterSilentHandler("/", []string{http.MethodGet}, handlers.HandleHealthCheck)
	app.RegisterNotFoundHandler()

	app.Run(*port)
}

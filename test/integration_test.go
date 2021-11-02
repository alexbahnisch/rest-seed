package test

import (
	"app/internal/handlers"
	"app/internal/server"
	"net/http"
	"os"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		app := server.New()

		app.RegisterHandler("/version", []string{http.MethodGet}, handlers.HandleVersion)
		app.RegisterSilentHandler("/", []string{http.MethodGet}, handlers.HandleHealthCheck)
		app.RegisterNotFoundHandler()

		wg.Done()
		app.Run(Port)
	}()

	wg.Wait()
	os.Exit(m.Run())
}

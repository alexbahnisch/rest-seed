package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func New() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

type Handler func(*http.Request) (int, []byte, error)

type NotFoundHandler struct{}

func (n NotFoundHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Printf("[%s] '%s': %d\n", request.Method, request.URL, 404)
	http.NotFound(response, request)
}

func (s *Server) RegisterNotFoundHandler() {
	s.router.NotFoundHandler = NotFoundHandler{}
}

func (s *Server) RegisterHandler(pattern string, methods []string, handler Handler) {
	s.router.HandleFunc(pattern, func(response http.ResponseWriter, request *http.Request) {
		statusCode, body, err := handler(request)
		if err != nil {
			log.Printf("[%s] '%s': %d\n", request.Method, request.URL, statusCode)
			http.Error(response, http.StatusText(statusCode), statusCode)
		}

		_, err = response.Write(body)
		if err != nil {
			log.Printf("[%s] '%s': %d\n", request.Method, request.URL, http.StatusInternalServerError)
			http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		log.Printf("[%s] '%s': %d\n", request.Method, request.URL, statusCode)
	}).Methods(methods...)
}

func (s *Server) RegisterSilentHandler(pattern string, methods []string, handler Handler) {
	s.router.HandleFunc(pattern, func(response http.ResponseWriter, request *http.Request) {
		statusCode, body, err := handler(request)
		if err != nil {
			http.Error(response, http.StatusText(statusCode), statusCode)
		}

		_, err = response.Write(body)
		if err != nil {
			http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}).Methods(methods...)
}

func (s *Server) Run(port int) {
	log.Printf("Server listing on port '%d' ...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), s.router))
}

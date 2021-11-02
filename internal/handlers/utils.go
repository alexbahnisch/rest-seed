package handlers

import "net/http"

func HandleHealthCheck(request *http.Request) (int, []byte, error) {
	return 200, nil, nil
}

func HandleVersion(request *http.Request) (int, []byte, error) {
	return 200, []byte("app@0.0.1"), nil
}

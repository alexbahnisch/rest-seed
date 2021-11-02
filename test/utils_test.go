package test

import (
	"testing"
)

func TestHealthCheck(t *testing.T) {
	if err := testGet("/", 200, nil); err != nil {
		t.Error(err)
	}
}

func TestVersionCheck(t *testing.T) {
	if err := testGet("/version", 200, []byte("app@0.0.1")); err != nil {
		t.Error(err)
	}
}

func TestNotFoundCheck(t *testing.T) {
	if err := testGet("/notfound", 404, []byte("404 page not found\n")); err != nil {
		t.Error(err)
	}
}

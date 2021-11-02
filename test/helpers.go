package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const Port = 9000

var baseUrl = fmt.Sprintf("http://localhost:%d", Port)

func testGet(route string, expectedStatusCode int, expectedBody []byte) error {
	resp, err := http.Get(baseUrl + route)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != expectedStatusCode {
		return StatusNotEqualError(route, expectedStatusCode, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if bytes.Compare(body, expectedBody) != 0 {
		return BodyNotEqualError(route, expectedBody, body)
	}

	return nil
}

func BodyNotEqualError(route string, expectedBody interface{}, actualBody interface{}) error {
	return fmt.Errorf("failed testing url '%s', expected response body: \n'%s'\n!= actual response body:\n'%s'", route, expectedBody, actualBody)
}

func StatusNotEqualError(route string, expectedStatusCode int, actualStatusCode int) error {
	return fmt.Errorf("failed testing url '%s', expected status code: '%d' != actual status code: '%d'", route, expectedStatusCode, actualStatusCode)
}

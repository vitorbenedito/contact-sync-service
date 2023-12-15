package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	GET = "GET"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestGetPaymentTypes(t *testing.T) {
	response := executeRequest(request(GET, "/health"))
	checkResponseCode(t, http.StatusOK, response.Code)
	resp, _ := responseToMap(response)
	if resp["status"] != "UP" {
		t.Errorf("Application is not UP")
	}

}

func responseToMap(response *httptest.ResponseRecorder) (map[string]string, error) {
	resp := map[string]string{}
	log.Println("Response: " + response.Body.String())
	var err = json.Unmarshal(response.Body.Bytes(), &resp)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}
	return resp, nil
}

func request(method string, endpoint string) *http.Request {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	return req
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

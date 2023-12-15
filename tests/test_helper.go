package tests

import (
	"contact-sync-service/api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseToMap(response *httptest.ResponseRecorder) (map[string]string, error) {
	resp := map[string]string{}
	log.Println("Response: " + response.Body.String())
	var err = json.Unmarshal(response.Body.Bytes(), &resp)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}
	return resp, nil
}

func Request(method string, endpoint string) *http.Request {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	return req
}

func ExecuteRequest(req *http.Request, a api.App) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

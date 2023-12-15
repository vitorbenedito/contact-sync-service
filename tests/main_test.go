package tests

import (
	"contact-sync-service/api"
	"net/http"
	"os"
	"testing"
)

const (
	GET = "GET"
)

var a api.App

func TestMain(m *testing.M) {
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestGetPaymentTypes(t *testing.T) {
	response := ExecuteRequest(Request(GET, "/health"), a)
	CheckResponseCode(t, http.StatusOK, response.Code)
	resp, _ := ResponseToMap(response)
	if resp["status"] != "UP" {
		t.Errorf("Application is not UP")
	}

}

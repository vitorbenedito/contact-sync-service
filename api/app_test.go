package api

import (
	"net/http"
	"os"
	"testing"
	"time"
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

func TetHealthCheck(t *testing.T) {
	response := ExecuteRequest(Request(GET, "/health"), a)
	CheckResponseCode(t, http.StatusOK, response.Code)
	resp, _ := ResponseToMap(response)
	if resp["status"] != "UP" {
		t.Errorf("Application is not UP")
	}

}

func TestSyncContacts(t *testing.T) {
	response := ExecuteRequest(Request(GET, "/contacts/sync"), a)
	CheckResponseCode(t, http.StatusOK, response.Code)
	resp, _ := ResponseToMap(response)
	if resp["syncedContacts"].(float64) != 24 {
		t.Errorf("Synced contacts not worked as expected")
	}
}

func TestSyncContactsAsync(t *testing.T) {
	response := ExecuteRequest(Request(GET, "/contacts/sync/async"), a)
	CheckResponseCode(t, http.StatusAccepted, response.Code)

	time.Sleep(3 * time.Second)

	responseGet := ExecuteRequest(Request(GET, "/contacts"), a)
	CheckResponseCode(t, http.StatusOK, responseGet.Code)
	resp, _ := ResponseToMap(responseGet)
	if resp["syncedContacts"].(float64) == 0 {
		t.Errorf("Synced contacts not worked as expected")
	}
}

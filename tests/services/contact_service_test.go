package tests

import (
	"contact-sync-service/services"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	os.Setenv("MAILCHIMP_API", "")
	code := m.Run()
	os.Exit(code)
}

func TestSyncContacts(t *testing.T) {

	services.SyncContacts()

	contacts := services.GetSyncedContatcs()

	if len(contacts) == 0 {
		t.Errorf("Contacts not found")
	}

}

func TestSyncContactsAsync(t *testing.T) {

	services.SyncContactsAsync()

	time.Sleep(3 * time.Second)

	contacts := services.GetSyncedContatcs()

	if len(contacts) == 0 {
		t.Errorf("Contacts not found")
	}
}

package services

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	os.Setenv("MAILCHIMP_API", "")
}

func TestSyncContacts(t *testing.T) {

	SyncContacts()

	contacts := GetSyncedContatcs()

	if len(contacts) == 0 {
		t.Errorf("Contacts not found")
	}

}

func TestSyncContactsAsync(t *testing.T) {

	SyncContactsAsync()

	time.Sleep(3 * time.Second)

	contacts := GetSyncedContatcs()

	if len(contacts) == 0 {
		t.Errorf("Contacts not found")
	}
}

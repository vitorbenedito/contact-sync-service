package services

import (
	"testing"
	"time"
)

func TestMain(m *testing.M) {
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

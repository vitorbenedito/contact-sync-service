package contact

import (
	"contact-sync-service/clients"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestGetContacts(t *testing.T) {

	contacts, _ := clients.GetContacts()
	if len(contacts) == 0 {
		t.Errorf("Contacts not found")
	}

}

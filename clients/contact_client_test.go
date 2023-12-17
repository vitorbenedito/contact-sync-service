package clients

import (
	"testing"
)

func TestGetContacts(t *testing.T) {
	contacts, _ := GetContacts()
	if len(contacts) == 0 {
		t.Errorf("Contacts not found")
	}

}

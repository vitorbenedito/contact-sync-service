package contact

import (
	"contact-sync-service/clients"
	"contact-sync-service/domains"
	"testing"
)

func TestSyncMembers(t *testing.T) {

	var contact = domains.Contact{
		Email:     "jhondutton@gmail.com",
		FirstName: "Jhon",
		LastName:  "Dutton",
	}

	member, _ := clients.SyncMembers(contact)
	if member.Email != contact.Email {
		t.Errorf("Error to sync member")
	}

}

package contact

import (
	"contact-sync-service/client"
	"contact-sync-service/domain"
	"testing"
)

func TestSyncMembers(t *testing.T) {

	contact := &domain.Contact{
		Email:     "jhondutton@gmail.com",
		FirstName: "Jhon",
		LastName:  "Dutton",
	}

	member, _ := client.SyncMembers(contact)
	if member.Email != contact.Email {
		t.Errorf("Error to sync member")
	}

}

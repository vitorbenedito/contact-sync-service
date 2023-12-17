package contact

import (
	"contact-sync-service/clients"
	"contact-sync-service/domains"
	"os"
	"testing"
)

func TestSyncMembers(t *testing.T) {

	os.Setenv("MAILCHIMP_API", "")

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

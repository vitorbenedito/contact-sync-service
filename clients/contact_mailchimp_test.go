package clients

import (
	"contact-sync-service/domains"
	"os"
	"testing"
)

func TestSyncMembers(t *testing.T) {

	os.Setenv("MAILCHIMP_API", "")

	var contact = domains.Contact{
		Email:     "manoelbandeira@gmail.com",
		FirstName: "Manoel",
		LastName:  "Bandeira",
	}

	var lists, _ = GetList()

	member, _ := SyncMembers(contact, lists[0].Id)
	if member.Email != contact.Email {
		t.Errorf("Error to sync member")
	}

}

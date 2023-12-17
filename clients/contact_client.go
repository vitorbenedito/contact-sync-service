package clients

import (
	"contact-sync-service/domains"
	"encoding/json"
	"log"
	"net/http"
)

func GetContacts() ([]domains.Contact, error) {

	r, err := http.Get("https://challenge.trio.dev/api/v1/contacts")
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Body.Close()

	var contacts []domains.Contact
	_ = json.NewDecoder(r.Body).Decode(&contacts)

	return contacts, nil
}

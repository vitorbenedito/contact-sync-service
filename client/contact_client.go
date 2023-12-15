package client

import (
	"contact-sync-service/domain"
	"encoding/json"
	"log"
	"net/http"
)

func GetContacts() ([]*domain.Contact, error) {

	r, err := http.Get("https://challenge.trio.dev/api/v1/contacts")
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Body.Close()

	var contacts []*domain.Contact
	_ = json.NewDecoder(r.Body).Decode(&contacts)

	return contacts, nil
}

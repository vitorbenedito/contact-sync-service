package services

import (
	"contact-sync-service/clients"
	"contact-sync-service/domains"
)

var syncedContacts = make(map[string]domains.Contact)

func SyncContactsAsync() chan int {

	r := make(chan int)

	go func() {
		SyncContacts()
		r <- 1
	}()

	return r

}

func SyncContacts() (map[string]domains.Contact, error) {

	var contacts, _ = clients.GetContacts()
	var synced = make(map[string]domains.Contact)

	for _, c := range contacts {
		clients.SyncMembers(c)
		syncedContacts[c.Email] = c
		synced[c.Email] = c
	}

	return synced, nil

}

func GetSyncedContatcs() map[string]domains.Contact {
	return syncedContacts
}

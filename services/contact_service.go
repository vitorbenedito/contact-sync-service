package services

import (
	"contact-sync-service/clients"
	"contact-sync-service/domains"
	"sync"
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

func SyncContactsParallel() (map[string]domains.Contact, error) {

	var contacts, _ = clients.GetContacts()
	var synced = make(map[string]domains.Contact)

	var wg sync.WaitGroup
	wg.Add(len(contacts))

	for _, c := range contacts {

		go func(contact domains.Contact) {
			defer wg.Done()
			clients.SyncMembers(contact)
			syncedContacts[contact.Email] = contact
			synced[contact.Email] = contact
		}(c)

	}

	wg.Wait()

	return synced, nil

}

func GetSyncedContatcs() map[string]domains.Contact {
	return syncedContacts
}

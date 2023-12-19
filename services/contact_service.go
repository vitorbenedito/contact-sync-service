package services

import (
	"contact-sync-service/clients"
	"contact-sync-service/domains"
	"contact-sync-service/jsons"
	"encoding/json"
	"errors"
	"log"
	"sync"
)

var syncedContacts = make(map[string]domains.Contact)

func CreateOrFindList() (domains.List, error) {

	var list domains.List

	err := json.Unmarshal([]byte(jsons.GetJson()), &list)
	if err != nil {
		return list, err
	}

	var lists, _ = clients.GetList()

	var listResp domains.List

	if len(lists) == 0 {
		listResp, _ := clients.CreateList(list)
		return *listResp, nil
	} else {
		for _, value := range lists {
			if value.Name == list.Name {
				listResp = value
				break
			}
		}

		if listResp == (domains.List{}) {
			return listResp, errors.New("List not found")
		}

		return listResp, nil

	}

}

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

	list, _ := CreateOrFindList()

	for _, c := range contacts {
		_, err := clients.SyncMembers(c, list.Id)
		if err != nil {
			syncedContacts[c.Email] = c
			synced[c.Email] = c
		}
	}

	return synced, nil

}

func SyncContactsParallel() (map[string]domains.Contact, error) {

	var contacts, _ = clients.GetContacts()
	synced := make(map[string]domains.Contact)

	// define the channel buffer to limit the requests simultaneously
	const max = 10
	queue := make(chan domains.Contact, max)

	wg := &sync.WaitGroup{}

	list, _ := CreateOrFindList()

	for i := 0; i < max; i++ {
		wg.Add(1)
		go worker(wg, queue, list.Id, synced)
	}

	for _, c := range contacts {
		queue <- c
	}

	close(queue)
	wg.Wait()

	return synced, nil

}

func worker(wg *sync.WaitGroup, queue chan domains.Contact, listId string, synced map[string]domains.Contact) {
	defer wg.Done()
	for contact := range queue {
		log.Println("Sync contact: " + contact.Email)
		_, err := clients.SyncMembers(contact, listId)
		if err != nil {
			log.Println("Error to sync conctact: " + contact.Email + ", error: " + err.Error())
		} else {
			syncedContacts[contact.Email] = contact
			synced[contact.Email] = contact
		}
	}
}

func GetSyncedContatcs() map[string]domains.Contact {
	return syncedContacts
}

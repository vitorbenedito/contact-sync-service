package client

import (
	"bytes"
	"contact-sync-service/domain"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SyncMembers(contact *domain.Contact) (*domain.Member, error) {

	member := &domain.Member{
		Email:  contact.Email,
		Status: "subscribed",
		Fields: &domain.Fields{
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
		},
	}

	body, err := json.Marshal(member)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	emailHash := md5.Sum([]byte(contact.Email))

	r, err := Put("https://us21.api.mailchimp.com/3.0/lists/de45421c0d/members/"+hex.EncodeToString(emailHash[:]), body)
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(b))
		return nil, errors.New("Error to sync members: " + string(b))
	}

	var memberResp *domain.Member
	_ = json.NewDecoder(r.Body).Decode(&memberResp)

	return memberResp, nil
}

func Put(url string, body []byte) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth("anystring", ""))
	c := &http.Client{}
	return c.Do(req)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
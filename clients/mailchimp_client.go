package clients

import (
	"bytes"
	"contact-sync-service/domains"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetList() ([]domains.List, error) {

	r, err := Request(http.MethodGet, "https://us21.api.mailchimp.com/3.0/lists", nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	listResp := map[string][]domains.List{}
	_ = json.NewDecoder(r.Body).Decode(&listResp)

	list := listResp["lists"]

	return list, nil

}

func CreateList(list domains.List) (*domains.List, error) {

	body, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r, err := Request(http.MethodPost, "https://us21.api.mailchimp.com/3.0/lists", body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(b))
		return nil, errors.New("Error to create list: " + string(b))
	}

	listResp := &domains.List{}
	_ = json.NewDecoder(r.Body).Decode(&listResp)

	return listResp, nil
}

func DeleteList(listId string) error {

	r, err := Request(http.MethodDelete, "https://us21.api.mailchimp.com/3.0/lists/"+listId, nil)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if r.StatusCode != http.StatusOK {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(b))
		return errors.New("Error to delete list: " + string(b))
	}

	return nil
}

func SyncMembers(contact domains.Contact, listId string) (*domains.Member, error) {

	member := &domains.Member{
		Email:  contact.Email,
		Status: "subscribed",
		Fields: &domains.Fields{
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

	r, err := Request(http.MethodPut, fmt.Sprintf("https://us21.api.mailchimp.com/3.0/lists/%s/members/%s", listId, hex.EncodeToString(emailHash[:])), body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
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

	memberResp := &domains.Member{}
	_ = json.NewDecoder(r.Body).Decode(&memberResp)

	return memberResp, nil
}

func Request(method string, url string, body []byte) (resp *http.Response, err error) {
	var req *http.Request
	if len(body) > 0 {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	req.Header.Set("Content-Type", "application/json")
	var apiKey = os.Getenv("MAILCHIMP_API")
	req.Header.Add("Authorization", "Basic "+basicAuth("anystring", apiKey))
	c := &http.Client{}
	return c.Do(req)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

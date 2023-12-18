# contact-sync-service

## About the project

Golang service to sync contacts from Contacts API and add members on a List in Mailchimp API

### Design

The project design documentation

* [Project Design](https://holly-card-34d.notion.site/Contacts-Sync-Technical-Design-d082fef3f0f24491908668392ddf41c3)

## Getting started

How deploy the application:

Configure MailChimp Key "MAILCHIMP_API" on api/.env file or define the Key on the variables

```
export MAILCHIMP_API=xyz
```

To deploy the project you can use the bash ./run on the project root

```
./run
```

You can deploy too using the go command

```
go run .
```

Here are the endpoints from the API

GET /contacts/sync

```
http://localhost:8080/contacts/sync
```

GET /contacts

```
http://localhost:8080/contacts
```

GET /contacts/sync/async

```
http://localhost:8080/contacts/sync/async
```

GET /contacts/sync/parallel

```
http://localhost:8080/contacts/sync/parallel
```
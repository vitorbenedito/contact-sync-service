package api

import (
	"contact-sync-service/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

type App struct {
	Router *gin.Engine
}

func (a *App) Initialize() {
	a.Router = gin.Default()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(a.Router.Run(addr))
}

func (a *App) initializeRoutes() {
	//public
	a.Router.GET("/health", Health)
	a.Router.GET("/contacts/sync", SyncContacts)
	a.Router.GET("/contacts/sync/parallel", SyncContactsParallel)
	a.Router.GET("/contacts/sync/async", SyncContactsAsync)
	a.Router.GET("/contacts", GetSyncedContacts)
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func SyncContactsAsync(c *gin.Context) {
	services.SyncContactsAsync()
	c.Status(http.StatusAccepted)
}

func SyncContacts(c *gin.Context) {
	contacts, _ := services.SyncContacts()

	c.JSON(http.StatusOK, gin.H{
		"syncedContacts": len(contacts),
		"contacts":       maps.Values(contacts),
	})
}

func SyncContactsParallel(c *gin.Context) {
	contacts, _ := services.SyncContactsParallel()

	c.JSON(http.StatusOK, gin.H{
		"syncedContacts": len(contacts),
		"contacts":       maps.Values(contacts),
	})
}

func GetSyncedContacts(c *gin.Context) {
	contacts := services.GetSyncedContatcs()

	c.JSON(http.StatusOK, gin.H{
		"syncedContacts": len(contacts),
		"contacts":       maps.Values(contacts),
	})
}

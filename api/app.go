package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

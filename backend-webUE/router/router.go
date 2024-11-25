package router

import (
	"backend-webUE/api"
	"backend-webUE/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ueProfileAPI *api.UeProfileAPI, serverConfig config.ServerConfig) *gin.Engine {

	// Initialize router
	router := gin.Default()

	ueProfileAPI.RegisterRoutes(router)

	return router
}

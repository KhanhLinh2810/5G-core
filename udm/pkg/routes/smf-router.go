package routes

import (
	"github.com/KhanhLinh2810/5G-core/udm/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func SMFRoutes(router *gin.Engine) {
	smfRouter := router.Group("/")
	{
		smfRouter.GET("/nudm-sdm/v2/:imsi/sm-data", controllers.GetSDMDetail)
		smfRouter.POST("/init-data/:rowDb", controllers.CreateSession)
		smfRouter.GET("/sessions", controllers.GetSessions)
	}
}

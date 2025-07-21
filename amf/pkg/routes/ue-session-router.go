package routes

import (
	"github.com/KhanhLinh2810/5G-core/amf/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func UESessionRoutes(router *gin.Engine) {
	ueSessionRouter := router.Group("/ue-session")
	{
		ueSessionRouter.GET("/", controllers.UECreateSession)
	}
}

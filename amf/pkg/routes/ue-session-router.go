package routes

import (
	"github.com/KhanhLinh2810/5G-core/amf/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func UESessionRoutes(router *gin.Engine) {
	ueSessionRouter := router.Group("/")
	{
		ueSessionRouter.GET("/ue-session", controllers.UECreateSession)
		ueSessionRouter.GET("/multi-ue-session", controllers.MultiUECreateSession)
		ueSessionRouter.POST("/namf-comm/v1/ue-context/:imsi/n1-n2-messages", controllers.N1N2MessageTransfer)
		ueSessionRouter.GET("/release-ue-session", controllers.ReleaseUECreateSession)
		ueSessionRouter.GET("/update-ue-session", controllers.UpdateUECreateSession)


	}
}

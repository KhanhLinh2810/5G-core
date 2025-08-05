package routes

import (
	"github.com/KhanhLinh2810/5G-core/smf/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func AMFRoutes(router *gin.Engine) {
	amfRouter := router.Group("/")
	{
		amfRouter.POST("/nsmf-pdusession/v1/sm-contexts", controllers.AMFCreateSession)
	}
}

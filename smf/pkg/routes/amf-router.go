package routes

import (
	"github.com/KhanhLinh2810/5G-core/smf/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func AMFRoutes(router *gin.Engine) {
	amfRouter := router.Group("/")
	{
		amfRouter.POST("/nsmf-pdusession/v1/sm-contexts", controllers.CreateSessionSaveInMap)
		amfRouter.POST("/nsmf-pdusession/v1/update-sm-contexts", controllers.UpdateSession)
		amfRouter.POST("/nsmf-pdusession/v1/release-sm-contexts", controllers.ReleaseSession)
	}
}

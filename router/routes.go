package router

import (
	"github.com/MarcosSmeets/Goportunities/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handler.GetListOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings", handler.DeleteOpeningHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
		v1.GET("/opening", handler.GetOpeningHandler)
	}
}

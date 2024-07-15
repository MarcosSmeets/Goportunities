package router

import (
	openings "github.com/MarcosSmeets/Goportunities/handler/openings"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", openings.GetListOpeningHandler)
		v1.POST("/openings", openings.CreateOpeningHandler)
		v1.DELETE("/openings", openings.DeleteOpeningHandler)
		v1.PUT("/openings", openings.UpdateOpeningHandler)
		v1.GET("/opening", openings.GetOpeningHandler)
	}
}

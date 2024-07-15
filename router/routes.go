package router

import (
	docs "github.com/MarcosSmeets/Goportunities/docs"
	"github.com/MarcosSmeets/Goportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/openings", handler.GetListOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings", handler.DeleteOpeningHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
		v1.GET("/opening", handler.GetOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

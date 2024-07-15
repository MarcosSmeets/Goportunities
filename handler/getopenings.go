package handler

import (
	"net/http"

	"github.com/MarcosSmeets/Goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	sendSuccess(ctx, "get-opening", opening)
}

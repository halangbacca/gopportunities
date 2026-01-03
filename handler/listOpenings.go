package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halangbacca/gopportunities/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	var openings []schemas.Opening

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "listOpenings", openings)
}

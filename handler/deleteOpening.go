package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halangbacca/gopportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "deleteOpening", opening)
}

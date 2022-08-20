package challcontrollers

import (
	"backend/controllers/middlewares.go"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type delReq struct {
	ID int `json:"id"`
}

func deleteChall(ctx *gin.Context) {
	var req delReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error binding JSON"})
		return
	}

	chall, err := models.GetChallenge(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting challenge"})
		return
	}

	err = chall.Delete()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error deleting challenge"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func RegisterDeleteChall(router *gin.RouterGroup) {
	router.POST("/delete", middlewares.EnsureAdmin(), deleteChall)
}

package challcontrollers

import (
	"backend/controllers/middlewares.go"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createChall(ctx *gin.Context) {
	var chall models.Challenge
	err := ctx.BindJSON(&chall)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error binding JSON"})
		return
	}

	err = chall.Save()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error saving challenge"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func RegisterCreateChall(router *gin.RouterGroup) {
	router.POST("/create", middlewares.EnsureAdmin(), createChall)
}

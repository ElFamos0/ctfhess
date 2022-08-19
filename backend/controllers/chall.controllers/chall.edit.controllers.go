package challcontrollers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func editChall(ctx *gin.Context) {
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

func RegisterEditChall(router *gin.RouterGroup) {
	router.POST("/edit", editChall)
}

package challcontrollers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func editChall(ctx *gin.Context) {
	var challSave models.Challenge
	err := ctx.BindJSON(&challSave)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error binding JSON"})
		return
	}

	challDB, err := models.GetChallenge(challSave.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting challenge"})
		return
	}
	challDB.Delete()

	challSave.ID = 0
	err = challSave.Save()
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

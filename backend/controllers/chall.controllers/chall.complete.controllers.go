package challcontrollers

import (
	"backend/controllers/middlewares.go"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type complReq struct {
	ID   int    `json:"id"`
	Flag string `json:"flag"`
}

func completeChall(ctx *gin.Context) {
	var req complReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error binding JSON"})
		return
	}

	chall, err := models.GetChallenge(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "Error getting challenge"})
		return
	}

	if chall.Flag != req.Flag {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "Flag is incorrect"})
		return
	}

	c := &models.Completion{
		ChallID: req.ID,
		UserID:  ctx.GetString("user_id"),
	}

	c.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func RegisterCompleteChall(router *gin.RouterGroup) {
	router.POST("/complete", middlewares.EnsureLoggedIn(), completeChall)
}

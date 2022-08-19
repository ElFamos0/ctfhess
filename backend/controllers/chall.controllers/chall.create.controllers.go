package challcontrollers

import (
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createChall(ctx *gin.Context) {
	// jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error binding JSON"})
	// 	return
	// }
	// fmt.Println(string(jsonData))

	var chall models.Challenge
	err := ctx.BindJSON(&chall)
	if err != nil {
		fmt.Println(err)
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
	router.POST("/create", createChall)
}

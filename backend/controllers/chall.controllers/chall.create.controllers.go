package challcontrollers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createChall(ctx *gin.Context) {
	user_id, _ := ctx.Get("user_id")

	user_id_int, _ := user_id.(int)

	user, err := services.GetUserByID(user_id_int)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
	}

	response := services.NewProfileDataResponse(user)

	ctx.JSON(http.StatusOK, response)
}

func RegisterCreateChall(router *gin.RouterGroup) {
	router.GET("/create", createChall)
}

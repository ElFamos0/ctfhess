package userscontrollers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func get_user_data(ctx *gin.Context) {
	user_id, _ := ctx.Get("user_id")

	user_id_int, _ := user_id.(int)

	user, err := services.GetUserByID(user_id_int)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
	}

	response := services.NewProfileDataResponse(user)

	ctx.JSON(http.StatusOK, response)
}

func Register_user_controllers_get(router *gin.RouterGroup) {
	router.GET("/data", get_user_data)
}

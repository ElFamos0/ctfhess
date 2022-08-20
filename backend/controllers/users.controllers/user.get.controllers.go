package userscontrollers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	u, err := ctx.Get("user")
	if !err {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
		return
	}
	user := u.(*models.User)
	response := user

	ctx.JSON(http.StatusOK, response)
}

func RegisterUserControllersGet(router *gin.RouterGroup) {
	router.GET("/data", GetUser)
}

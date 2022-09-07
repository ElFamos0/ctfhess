package userscontrollers

import (
	"backend/controllers/middlewares.go"
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

func GetAllUsers(ctx *gin.Context) {
	users := models.GetUserList()
	for _, u := range users {
		u.Points = u.GetPoints()
	}
	ctx.JSON(http.StatusOK, users)
}

func RegisterUserControllersGet(router *gin.RouterGroup) {
	router.GET("/data", GetUser)
	router.GET("/list", middlewares.EnsureAdmin(), GetAllUsers)
}

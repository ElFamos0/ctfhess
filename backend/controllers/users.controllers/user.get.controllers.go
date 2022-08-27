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

func GetAllAdmins(ctx *gin.Context) {
	admins := models.GetAdminList()
	ctx.JSON(http.StatusOK, admins)
}

func GetAllUsers(ctx *gin.Context) {
	users := models.GetUserList()
	ctx.JSON(http.StatusOK, users)
}

func RegisterUserControllersGet(router *gin.RouterGroup) {
	router.GET("/data", GetUser)
	router.GET("/admins", middlewares.EnsureAdmin(), GetAllAdmins)
	router.GET("/list", middlewares.EnsureAdmin(), GetAllUsers)
}

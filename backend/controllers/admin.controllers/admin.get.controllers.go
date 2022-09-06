package admincontrollers

import (
	"backend/controllers/middlewares.go"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAdmins(ctx *gin.Context) {
	admins := models.GetAdminList()
	ctx.JSON(http.StatusOK, admins)
}

func RegisterAdminGet(router *gin.RouterGroup) {
	router.GET("/get", middlewares.EnsureAdmin(), GetAllAdmins)
}

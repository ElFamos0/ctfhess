package admincontrollers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAdmins(ctx *gin.Context) {
	admins := models.GetAdminList()
	for _, u := range admins {
		u.Points = u.GetPoints()
	}
	ctx.JSON(http.StatusOK, admins)
}

func RegisterAdminGet(router *gin.RouterGroup) {
	router.GET("/get", GetAllAdmins)
}

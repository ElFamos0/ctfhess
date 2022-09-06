package admincontrollers

import (
	"backend/controllers/middlewares.go"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeAdmin(ctx *gin.Context) {
	var req struct {
		ID string `json:"ID"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request."})
		return
	}

	// Get the user from the database
	user := models.GetUserByID(req.ID)
	// Set the user type to admin
	user.MakeAdmin()
	ctx.JSON(http.StatusOK, gin.H{"message": "User is now an admin."})
}

func MakeNormal(ctx *gin.Context) {
	var req struct {
		ID string `json:"ID"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request."})
		return
	}

	// Get the user from the database
	user := models.GetUserByID(req.ID)
	// Set the user type to admin
	user.MakeNormal()
	ctx.JSON(http.StatusOK, gin.H{"message": "User is not an admin anymore."})
}

func RegisterAdminChange(router *gin.RouterGroup) {
	router.POST("/add", middlewares.EnsureAdmin(), MakeAdmin)
	router.POST("/remove", middlewares.EnsureAdmin(), MakeNormal)
}

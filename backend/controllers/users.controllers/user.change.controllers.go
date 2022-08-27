package userscontrollers

import (
	"backend/controllers/middlewares.go"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAdmin(ctx *gin.Context) {
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

func RegisterUserControllersAdmin(router *gin.RouterGroup) {
	router.POST("/makeadmin", middlewares.EnsureAdmin(), UserAdmin)
}

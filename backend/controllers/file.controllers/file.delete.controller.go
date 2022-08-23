package filecontrollers

import (
	"backend/controllers/middlewares.go"
	"backend/db"
	"backend/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func deleteFile(ctx *gin.Context) {
	fileName := ctx.Param("name")

	err := os.Remove("data/" + fileName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error deleting file"})
		return
	}

	db.DB.Delete(models.File{}, "uri = ?", fileName)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func RegisterDeleteFile(router *gin.RouterGroup) {
	router.GET("/delete/:name", middlewares.EnsureAdmin(), deleteFile)
}

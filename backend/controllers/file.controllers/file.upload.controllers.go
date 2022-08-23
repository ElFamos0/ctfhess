package filecontrollers

import (
	"backend/controllers/middlewares.go"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createFile(ctx *gin.Context) {
	form, err := ctx.MultipartForm()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting form"})
		return
	}
	log.Println(form.File)
	files := form.File["file"]
	log.Println(files)
	for _, file := range files {
		log.Println(file.Filename)
		err := ctx.SaveUploadedFile(file, "data/"+file.Filename)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
			return
		}

	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})

}

func RegisterCreateFile(router *gin.RouterGroup) {
	router.POST("/upload", middlewares.EnsureAdmin(), createFile)
}

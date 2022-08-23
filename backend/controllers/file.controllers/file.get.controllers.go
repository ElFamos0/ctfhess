package filecontrollers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func getFile(ctx *gin.Context) {
	fileName := ctx.Param("name")

	data, err := os.ReadFile("data/" + fileName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
		return
	}

	// We reply image for images type and octet stream for other types
	if strings.HasSuffix(fileName, "png") {
		ctx.Writer.Header().Set("Content-Type", "image/png")
	} else if strings.HasSuffix(fileName, "jpg") {
		ctx.Writer.Header().Set("Content-Type", "image/jpg")
	} else if strings.HasSuffix(fileName, "jpeg") {
		ctx.Writer.Header().Set("Content-Type", "image/jpeg")
	} else {
		ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
	}
	ctx.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	ctx.Writer.Write(data)
}

func RegisterGetFile(router *gin.RouterGroup) {
	router.GET("/get/:name", getFile)
}

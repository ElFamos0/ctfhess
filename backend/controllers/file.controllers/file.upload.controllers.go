package filecontrollers

import (
	"backend/controllers/middlewares.go"
	"backend/db"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func createFile(ctx *gin.Context) {
	form, err := ctx.MultipartForm()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting form"})
		return
	}

	// Get the last chall id
	var challID int
	tx := db.DB.Raw("SELECT id FROM challenges ORDER BY id DESC LIMIT 1").Scan(&challID)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting last chall id"})
		return
	}

	// Get all the pages of the chall
	var pages []models.ChallengePage
	tx = db.DB.Where("challenge_id = ?", challID).Find(&pages)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting pages"})
		return
	}

	for i, page := range pages {
		files := form.File[fmt.Sprintf("files[%d]", i)]
		for _, file := range files {
			newName, err := uuid.NewV4()
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error creating uuid"})
				return
			}

			filename := fmt.Sprintf("%s_%s", newName.String(), file.Filename)

			err = ctx.SaveUploadedFile(file, "data/"+filename)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
				return
			}

			newFile := models.File{
				PageID: page.ID,
				Name:   file.Filename,
				URI:    filename,
			}

			newFile.Save()
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func RegisterCreateFile(router *gin.RouterGroup) {
	router.POST("/upload", middlewares.EnsureAdmin(), createFile)
}

package filecontrollers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	os.Mkdir("data", 0755)
}

func RegisterFileRoutes(router *gin.RouterGroup) {
	RegisterCreateFile(router)
	RegisterGetFile(router)
	RegisterDeleteFile(router)
	RegisterReuploadFiles(router)
}

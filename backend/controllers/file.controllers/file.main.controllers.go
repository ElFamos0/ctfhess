package filecontrollers

import "github.com/gin-gonic/gin"

func RegisterFileRoutes(router *gin.RouterGroup) {
	RegisterCreateFile(router)
	RegisterGetFile(router)
	RegisterDeleteFile(router)
	RegisterReuploadFiles(router)
}

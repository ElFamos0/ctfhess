package challcontrollers

import "github.com/gin-gonic/gin"

func RegisterChallRoutes(router *gin.RouterGroup) {

	RegisterCreateChall(router)
	RegisterGetAllChall(router)
	RegisterDeleteChall(router)
	RegisterEditChall(router)
	RegisterGetAllChallAdmin(router)
}

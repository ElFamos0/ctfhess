package admincontrollers

import "github.com/gin-gonic/gin"

func RegisterAdminRoutes(router *gin.RouterGroup) {

	RegisterAdminGet(router)
	RegisterAdminChange(router)
}

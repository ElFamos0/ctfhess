package userscontrollers

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(router *gin.RouterGroup) {

	RegisterUserControllersGet(router)
	RegisterScoreboardControllersGet(router)
}

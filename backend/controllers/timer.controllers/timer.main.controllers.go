package timercontrollers

import "github.com/gin-gonic/gin"

func RegisterTimerRoutes(router *gin.RouterGroup) {

	RegisterTimerGet(router)
}

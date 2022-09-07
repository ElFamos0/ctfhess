package timercontrollers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AbsTimer struct {
	TimeLeft int `json:"timeleft"`
}

func GetTimerAbs(ctx *gin.Context) {
	hours := 24 - int(time.Since(parsedTime).Hours())%24
	minutes := 60 - int(time.Since(parsedTime).Minutes())%60
	seconds := 60 - int(time.Since(parsedTime).Seconds())%60

	ctx.JSON(http.StatusOK, &AbsTimer{
		TimeLeft: hours*3600 + minutes*60 + seconds,
	})
}

func RegisterTimerGetAbs(router *gin.RouterGroup) {
	router.GET("/getAbs", GetTimerAbs)
}

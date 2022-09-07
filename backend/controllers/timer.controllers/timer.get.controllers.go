package timercontrollers

import (
	"backend/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	startDateFormat = config.Getenv("CHALL_START_DATE", "2022-09-06T06:00:00")
	parsedTime      time.Time
)

type Timer struct {
	TimeElapsed       int `json:"timer"`
	TimeBeforeNextDay int `json:"timeleft"`
}

func init() {
	var err error
	parsedTime, err = time.Parse("2006-01-02T15:04:05", startDateFormat)
	if err != nil {
		panic(err)
	}
}

func GetTimer(ctx *gin.Context) {
	hoursLeft := 24 - int(time.Since(parsedTime).Hours())%24
	minutesLeft := 60 - int(time.Since(parsedTime).Minutes())%60
	secondsLeft := 60 - int(time.Since(parsedTime).Seconds())%60

	ctx.JSON(http.StatusOK, &Timer{
		TimeElapsed:       int(time.Since(parsedTime).Seconds()),
		TimeBeforeNextDay: hoursLeft*3600 + minutesLeft*60 + secondsLeft,
	})
}

func RegisterTimerGet(router *gin.RouterGroup) {
	router.GET("/get", GetTimer)
}

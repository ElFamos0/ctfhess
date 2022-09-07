package timercontrollers

import (
	"backend/config"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	startDateFormat = config.Getenv("CHALL_START_DATE", "2022-09-06T06:00:00")
	parsedTime      time.Time
)

type Timer struct {
	TimeElapsed       string `json:"timer"`
	TimeBeforeNextDay string `json:"timeleft"`
}

func init() {
	var err error
	parsedTime, err = time.Parse("2006-01-02T15:04:05", startDateFormat)
	if err != nil {
		panic(err)
	}
}

func GetTimer(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &Timer{
		TimeElapsed:       fmt.Sprintf("%d jours %d heures %d minutes %d secondes", int(time.Since(parsedTime).Hours()/24), int(time.Since(parsedTime).Hours())%24, int(time.Since(parsedTime).Minutes())%60, int(time.Since(parsedTime).Seconds())%60),
		TimeBeforeNextDay: fmt.Sprintf("%d heures %d minutes %d secondes", 24-int(time.Since(parsedTime).Hours())%24, 60-int(time.Since(parsedTime).Minutes())%60, 60-int(time.Since(parsedTime).Seconds())%60),
	})
}

func RegisterTimerGet(router *gin.RouterGroup) {
	router.GET("/get", GetTimer)
}

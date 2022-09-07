package challcontrollers

import (
	"backend/config"
	"backend/db"
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	maxday          = config.Getenvint("CHALL_DURATION_DAY", 31)
	startDateFormat = config.Getenv("CHALL_START_DATE", "2022-09-06T06:00:00")
	parsedTime      time.Time

	fakeChall = models.Challenge{
		Fake: true,
	}
)

func init() {
	var err error
	parsedTime, err = time.Parse("2006-01-02T15:04:05", startDateFormat)
	if err != nil {
		panic(err)
	}
}

func currentDay() int {
	return int(time.Since(parsedTime).Hours() / 24)
}

func getAllChall(ctx *gin.Context) {
	var logged bool
	var admin bool

	logged = ctx.GetBool("is_logged_in")
	if logged {
		u, found := ctx.Get("user")
		if found {
			if u.(*models.User).Type == models.AdminUser {
				admin = true
			}
		}
	}

	c, err := models.GetAllChallenges()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting challenges"})
		return
	}

	var challenges [][]*models.Challenge
	for day := 0; day < maxday; day++ {
		var L []*models.Challenge
		for _, c := range c {
			c.Flag = ""
			if c.OpensAt == day {
				if day > currentDay() && !admin {
					chall := fakeChall
					chall.OpensAt = day
					L = append(L, &chall)
					continue
				}
				L = append(L, c)
				// Fill completions
				db.DB.Model(&models.Completion{}).Where("chall_id = ?", c.ID).Count(&c.Completions)
			}
		}
		challenges = append(challenges, L)
	}

	// remove last empty days
	for day := maxday - 1; day >= 0; day-- {
		if len(challenges[day]) == 0 {
			challenges = challenges[:day]
		} else {
			break
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":    true,
		"challenges": challenges,
	})
}

func RegisterGetAllChall(router *gin.RouterGroup) {
	router.GET("/getall", getAllChall)
}

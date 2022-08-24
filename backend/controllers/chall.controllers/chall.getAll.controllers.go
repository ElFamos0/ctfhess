package challcontrollers

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MAXDAY = 31
)

var (
	currentDay = 1
	fakeChall  = &models.Challenge{
		Fake: true,
	}
)

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
	for day := 0; day < MAXDAY; day++ {
		var L []*models.Challenge
		for _, c := range c {
			c.Flag = ""
			if c.OpensAt == day {
				if day > currentDay && !admin {
					L = append(L, fakeChall)
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
	for day := MAXDAY - 1; day >= 0; day-- {
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

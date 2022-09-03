package challcontrollers

import (
	"backend/controllers/middlewares.go"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllChallAdmin(ctx *gin.Context) {
	c, err := models.GetAllChallenges()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting challenges"})
		return
	}

	var challenges [][]*models.Challenge
	for day := 0; day < maxday; day++ {
		var L []*models.Challenge
		for _, c := range c {
			if c.OpensAt == day {
				L = append(L, c)
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

func RegisterGetAllChallAdmin(router *gin.RouterGroup) {
	router.GET("/getalladmin", middlewares.EnsureAdmin(), getAllChallAdmin)
}

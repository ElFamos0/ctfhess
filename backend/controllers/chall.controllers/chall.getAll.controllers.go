package challcontrollers

import (
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MAXDAY = 31
)

func getAllChall(ctx *gin.Context) {
	c, err := models.GetAllChallenges()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting challenges"})
		return
	}

	var challenges [][]*models.Challenge
	for day := 0; day < MAXDAY; day++ {
		var L []*models.Challenge
		for _, c := range c {
			if c.OpensAt == day {
				L = append(L, c)
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
	fmt.Println(challenges[0][0].Pages)

	ctx.JSON(http.StatusOK, gin.H{
		"success":    true,
		"challenges": challenges,
	})
}

func RegisterGetAllChall(router *gin.RouterGroup) {
	router.GET("/getall", getAllChall)
}

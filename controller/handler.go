package controller

import (
	"github.com/gin-gonic/gin"
	"mangtas-test/helper"
	"strings"
)

func MostUsedWord() gin.HandlerFunc {
	return func(c *gin.Context) {
		content := &struct {
			Text string `json:"text"`
		}{}

		if err := c.BindJSON(content); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		result := helper.Counter(strings.TrimSpace(content.Text))
		c.JSON(200, gin.H{"Top 10 most used words": result})
	}
}

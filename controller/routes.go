package controller

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/", MostUsedWord())
}

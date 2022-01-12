package main

import (
	"github.com/gin-gonic/gin"
	"mangtas-test/controller"
)

func main() {
	r := gin.Default()
	controller.Routes(r)

	if err := r.Run(":8080"); err != nil {
		panic(err.Error())
	}

}

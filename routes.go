package main

import "github.com/gin-gonic/gin"

func initializeRoutes(r *gin.Engine) {
	r.GET("/", index)

}

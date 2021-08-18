package main

import "github.com/gin-gonic/gin"

func initializeRoutes(r *gin.Engine) {
	r.GET("/", index)

	// API
	r.GET("users/", getUsers)
	r.GET("users/:id")
	r.POST("users/", postUsers)
	r.PUT("users/:id")
	r.DELETE("users/:id")
}

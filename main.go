package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	initializeRoutes(router)
	err := router.Run()
	if err != nil {
		return
	}
}

func index(conn *gin.Context) {
	conn.HTML(http.StatusOK, "index.html", gin.H{"title": "Home page"})
}

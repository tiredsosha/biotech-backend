package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func routes() {
	router.GET("/ping", ping)
	router.GET("/", index)
	router.POST("/themes", themeSender)
}

func Start() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Use(cors.Default())

	routes()

	router.Run()
	fmt.Println("http is started")
}

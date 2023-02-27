package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func routes() {
	router.POST("/museumpower", museumPower)
	router.POST("/museumsound", museumSound)
	router.POST("/zonespower", zonesPower)
	router.POST("/zonessound", zonesSound)
	router.POST("/zonescontent", zonesContent)
	router.POST("/zonesballs", zonesBalls)
	router.POST("/zonesled", zonesLed)
	router.GET("/ping", ping)
	router.GET("/status", status)

	// V1 Group
	// museum := router.Group("/v1/museum").Use(myCors())
	// {
	// 	museum.POST("/power", museumPower)
	// 	museum.POST("/sound", museumSound)
	// }

	// zones := router.Group("/v1/zones").Use(myCors())
	// {
	// 	zones.POST("/power", zonesPower)
	// 	zones.POST("/sound", zonesSound)
	// 	zones.POST("/content", zonesContent)
	// 	zones.POST("/balls", zonesBalls)
	// 	zones.POST("/led", zonesLed)
	// }

	// common := router.Group("/v1").Use(myCors())
	// {
	// 	common.GET("/ping", ping)
	// 	common.GET("/status", status)
	// }
}

func Start() {
	// gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router := gin.Default()
	router.Use(myCors())

	routes()
	router.Run()

	fmt.Println("http is started")
}

func myCors() gin.HandlerFunc {
	newCors := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://127.0.0.1"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
	return newCors
}

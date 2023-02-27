package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Museum handlers
func museumPower(conn *gin.Context) {
	var req Basic
	conn.BindJSON(&req)
	fmt.Println(req)

	conn.Data(http.StatusOK, "application/text", []byte("OK"))
}

func museumSound(conn *gin.Context) {
	var req Basic
	conn.BindJSON(&req)
	fmt.Println(req)

	conn.Data(http.StatusOK, "application/text", []byte("OK"))
}

// Zones handlers
func zonesPower(conn *gin.Context) {
	var req Advansed
	conn.BindJSON(&req)
	fmt.Println(req)

	conn.Data(http.StatusOK, "application/text", []byte("OK"))
}

func zonesSound(conn *gin.Context) {
	var req Advansed
	conn.BindJSON(&req)
	fmt.Println(req)

	conn.Data(http.StatusOK, "application/text", []byte("OK"))
}

func zonesContent(conn *gin.Context) {
	var req Basic
	conn.BindJSON(&req)
	fmt.Println(req)

	conn.Data(http.StatusOK, "application/text", []byte("OK"))
}

func zonesBalls(conn *gin.Context) {
	var req Basic
	conn.BindJSON(&req)
	fmt.Println(req)

	conn.JSON(200, gin.H{
		"Status": "ok",
	})

	// conn.Data(http.StatusOK, "application/text", []byte("OK"))
}

func zonesLed(conn *gin.Context) {
	var req Advansed
	conn.BindJSON(&req)
	fmt.Println(req)

	conn.Data(http.StatusOK, "application/text", []byte("OK"))
}

// Status
func status(conn *gin.Context) {
	conn.Data(http.StatusOK, "application/text", []byte("ok"))
}

// Ping
func ping(conn *gin.Context) {
	conn.Data(http.StatusOK, "application/text", []byte("pong"))
}

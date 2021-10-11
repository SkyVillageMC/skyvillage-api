package main

import (
	"log"

	"github.com/SkyVillageMc/skyvillage-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting...")
	r := gin.Default()

	routes.InitAuth(r.Group("/auth"))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(":3000")
}

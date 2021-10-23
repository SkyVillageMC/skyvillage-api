package main

import (
	"log"

	"github.com/SkyVillageMc/skyvillage-api/database"
	"github.com/SkyVillageMc/skyvillage-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting...")
	r := gin.Default()

	database.Init()
	defer func() {
		if err := database.DB.Prisma.Disconnect(); err != nil {
			//Pretty sure this wouldn't happen anytime soon(maybe because of sqlite)
			log.Fatalf("Error diconnecting from the db\n%s\n", err.Error())
		}
	}()

	routes.InitWs(r)
	routes.InitPresence(r.Group("/presence"))
	routes.InitUsers(r.Group("/user"))
	routes.InitAuth(r.Group("/auth"))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(":3000")
}

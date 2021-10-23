package routes

import (
	"context"

	"github.com/SkyVillageMc/skyvillage-api/database"
	"github.com/SkyVillageMc/skyvillage-api/db"
	"github.com/gin-gonic/gin"
)

func InitAuth(r *gin.RouterGroup) {

	r.PUT("/register", func(c *gin.Context) {
		var data struct {
			Name string `json:"username" binding:"required"`
		}

		if c.BindJSON(&data) == nil {
			database.DB.User.CreateOne(
				db.User.Name.Set(data.Name),
				db.User.Presence.Link(
					db.Presence.PartyID.Equals("asd"),
				),
				db.User.PresenceID.Set(0),
				db.User.Party.Link(
					db.Party.ID.Equals("asd"),
				),
			).Exec(context.TODO())
			c.JSON(200, gin.H{
				"status": "success",
			})
			return
		}
		c.JSON(400, gin.H{
			"error": "missing data",
		})
	})
}

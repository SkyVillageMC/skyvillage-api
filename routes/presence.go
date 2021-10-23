package routes

import (
	"context"
	"errors"
	"log"

	"github.com/SkyVillageMc/skyvillage-api/database"
	"github.com/SkyVillageMc/skyvillage-api/db"
	"github.com/SkyVillageMc/skyvillage-api/models"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func InitPresence(r *gin.RouterGroup) {

	r.GET("/:id", func(c *gin.Context) {
		presence, err := database.DB.Presence.FindFirst(
			db.Presence.UserID.Equals(c.Param("id")),
		).Exec(ctx)

		if err != nil {
			if errors.Is(err, db.ErrNotFound) {
				c.JSON(404, gin.H{
					"error": "no presence found",
				})
				return
			}
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, *presence)
	})

	r.PATCH("/:id", func(c *gin.Context) {

		var data models.Presence

		if err := c.BindJSON(&data); err != nil {
			log.Println(err.Error())
			c.JSON(400, gin.H{
				"error": "incomplete data",
			})
			return
		}

		_, err := database.DB.Presence.UpsertOne(db.Presence.UserID.Equals(c.Param("id"))).Update(
			db.Presence.UserID.Set(c.Param("id")),
			db.Presence.User.Link(
				db.User.ID.Equals(c.Param("id")),
			),
			db.Presence.State.Set(data.State),
			db.Presence.LargeImageKey.Set(data.LargeImageKey),
			db.Presence.SmallImageKey.Set(data.SmallImageKey),
			db.Presence.StartTime.Set(db.BigInt(data.StartTime)),
			db.Presence.EndTime.Set(db.BigInt(data.EndTime)),
			db.Presence.Details.Set(data.Details),
			db.Presence.LargeImageText.Set(data.LargeImageText),
			db.Presence.SmallImageText.Set(data.SmallImageText),
			db.Presence.PartyID.Set(data.PartyId),
			db.Presence.Party.Link(
				db.Party.ID.Equals(data.PartyId),
			),
		).Exec(ctx)

		presence, _ := database.DB.Presence.FindFirst(
			db.Presence.UserID.Equals(c.Param("id")),
		).Exec(ctx)

		SendToUser(c.Param("id"), gin.H{
			"event": "presence-update",
			"data":  presence,
		})

		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, db.ErrNotFound) {
				c.JSON(404, gin.H{
					"error": "no presence found",
				})
				return
			}
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "success",
		})
	})
}

package routes

import (
	"errors"

	"github.com/SkyVillageMc/skyvillage-api/database"
	"github.com/SkyVillageMc/skyvillage-api/db"
	"github.com/gin-gonic/gin"
)

func InitParties(r *gin.RouterGroup) {
	r.GET("/:id", func(c *gin.Context) {
		party, err := database.DB.Party.FindFirst(
			db.Party.ID.Equals(c.Param("id")),
		).Exec(ctx)

		if err != nil {
			if errors.Is(err, db.ErrNotFound) {
				c.JSON(404, gin.H{
					"error": "no user found",
				})
				return
			}
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, *party)
	})
}

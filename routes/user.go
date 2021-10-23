package routes

import (
	"errors"

	"github.com/SkyVillageMc/skyvillage-api/database"
	"github.com/SkyVillageMc/skyvillage-api/db"
	"github.com/gin-gonic/gin"
)

func InitUsers(r *gin.RouterGroup) {
	r.GET("/:name", func(c *gin.Context) {
		user, err := database.DB.User.FindFirst(
			db.User.Name.Equals(c.Param("name")),
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

		c.JSON(200, *user)
	})
}

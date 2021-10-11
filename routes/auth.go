package routes

import "github.com/gin-gonic/gin"

func InitAuth(r *gin.RouterGroup) {

	r.PUT("/register", func(c *gin.Context) {
		var data struct {
			Name string `json:"username" binding:"required"`
		}

		if c.BindJSON(&data) == nil {
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

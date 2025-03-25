package routers

import (
	"github.com/gin-gonic/gin"
)

func NewHealthRouter(router *gin.Engine) {

	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
	}
}

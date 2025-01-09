package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Engine = gin.Default()

func InitRouter() {

	Engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Title": "TaskManagerWebb"})
	})

	Engine.Run(":9876")
}

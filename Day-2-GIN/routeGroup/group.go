package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api:=r.Group("/api")

	api.GET("/hello",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"hello from APi !",
		})
	})
	api.POST("/user",func(c *gin.Context) {
		var body map[string]interface{}
		c.BindJSON(&body)

		c.JSON(http.StatusOK,gin.H{
			"status":"user received",
			"data":body,
		})
	})
	r.Run(":8080")
}
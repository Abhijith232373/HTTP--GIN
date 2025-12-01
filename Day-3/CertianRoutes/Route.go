package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const USERNAME = "admin"
const PASSWORD = "12345"

func main() {
	r := gin.Default()

	r.POST("/login",func (c *gin.Context)  {
		var body struct{
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err:=c.BindJSON(&body);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"invalid request"})
		return
		}
		if body.Username==USERNAME&&body.Password==PASSWORD{
			c.SetCookie("session","loggedin",3600,"/","localhost",false,true)
			c.JSON(http.StatusOK,gin.H{"message":"login successful"})
		}else{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"invalid credentials"})
		}
	})

authMiddlware:=func (c *gin.Context){
	cookie,err:=c.Cookie("session")
	if err!=nil||cookie!="loggedin"{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":"plase login to access this route",
		})
		c.Abort()
		return
	}
	c.Next()
}


Protected:=r.Group("/protected")
Protected.Use(authMiddlware)

Protected.GET("/profile",func(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"username":"admin",
		"role":"user",
	})
})
r.POST("/logout",func(c *gin.Context) {
	c.SetCookie("session","",-1,"/","localhost",false,true)
	c.JSON(http.StatusOK,gin.H{"message":"logged out"})
})
r.Run(":8080")
}
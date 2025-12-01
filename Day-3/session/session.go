package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const USERNAME = "admin"
const PASSWORD = "123456"

func main() {
	r := gin.Default()

	r.POST("/login",func (c *gin.Context)  {
		var body struct{
			Username string `json:"username"`
			Password string `json:"password"`
 		}
		if err:=c.BindJSON(&body);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"invaild requst"})
			return
		}
		if body.Username==USERNAME&&body.Password==PASSWORD{
			c.SetCookie("session","loggedin",3600,"/","localhostt",false,true)
			c.JSON(http.StatusOK,gin.H{
				"message":"login succesful",
			})

		}else{
			c.JSON(http.StatusUnauthorized,gin.H{
				"error":"invaild username or password",
			})
		}
	})

	r.GET("/home",func(c *gin.Context) {
		session,err:=c.Cookie("session")
		if err !=nil||session!="loggedin"{
			c.JSON(http.StatusUnauthorized,gin.H{
				"error":"you must login first",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"message ": "Welcome ! you are logged in.",
		})
	})

	r.POST("/logout",func (c *gin.Context)  {
		c.SetCookie("session","",-1,"/","localhost",false,true)
		c.JSON(http.StatusOK,gin.H{
			"message":"logged out",
		})
	})
	r.Run(":8080")
}
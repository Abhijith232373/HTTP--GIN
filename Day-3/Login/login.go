package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
const USERNAME ="Abhijith"
const PASSWORD ="123345"

type LogonRequst struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.POST("/login",func (c *gin.Context)  {
		var req LogonRequst

		if err:=c.BindJSON(&req);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"invaild requst",
			})
			return
		}
		if req.Username==USERNAME&&req.Password==PASSWORD{
			c.JSON(http.StatusOK,gin.H{
				"message":"login successful",
			})

		}else{
			c.JSON(http.StatusUnauthorized,gin.H{
				"error":"invaild username or password",
			})
		}
	})
	r.Run(":8080")
}
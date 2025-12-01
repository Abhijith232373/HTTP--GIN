package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const USERNAME = "admin"
const PASSWORD = "12345"

func LoginRequsts() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("incoming Requst:",c.Request.Method,c.Request.URL.Path)

		c.Next()
		fmt.Println("completeed with status:",c.Writer.Status())
	}
}
func main(){
	r:=gin.Default()

	authRoute:=r.Group("/")
	authRoute.Use(LoginRequsts())
	
	authRoute.POST("/login",func (c *gin.Context){
		var body struct{
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err:=c.BindJSON(&body);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"invaild requst"})
			return
		}
		if body.Username==USERNAME&&body.Password==PASSWORD{
			c.JSON(http.StatusOK,gin.H{"message ":"login successful"})
		}else{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"invaild username or password"})
		}
	})
	authRoute.POST("/logout",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"logged out"})
	})
	r.Run(":8080")
}
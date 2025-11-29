package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Abhijith", Age: 20},
	{ID: 2, Name: "Appu", Age: 20},
}

func main() {
	r := gin.Default()

	api:=r.Group("/api")
	{
		api.GET("/users",getUsers)
		api.POST("/users",createUsers)
	}
	r.Run(":8080")

}
func getUsers(c *gin.Context){
	c.JSON(http.StatusOK,users)
}

func createUsers(c *gin.Context){
	var input struct{
		Name string `json:"name"`
		Age int `json:"age"`
	}
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid JSON"})
		return
	}
	if input.Name==""||input.Age<=0{
		c.JSON(http.StatusBadRequest,gin.H{"error":"name and age are requird"})
		return
	}
	newID:=len(users)+1
user:=User{
	ID: newID,
	Name:input.Name,
	Age: input.Age,
}
users=append(users,user)
c.JSON(http.StatusCreated,user)
}

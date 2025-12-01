package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var sessions = map[string]string{} 

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		println("Request:", method, path)
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_id")
		if err != nil || sessions[cookie] == "" {
			c.String(http.StatusUnauthorized, "Unauthorized. Please login.")
			c.Abort()
			return
		}
		c.Set("user", sessions[cookie])
		c.Next()
	}
}
func Login(c *gin.Context) {
	username := c.Query("user")

	if username == "" {
		c.String(400, "Username required")
		return
	}


	sessionID := username + "_123"
	sessions[sessionID] = username

	c.SetCookie("session_id", sessionID, 3600, "/", "", false, true)

	c.String(200, "Login successful for user: %s", username)
}

func Dashboard(c *gin.Context) {
	user, _ := c.Get("user")
	c.String(200, "Welcome to your dashboard, %s!", user)
}

func Logout(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err == nil {
		delete(sessions, cookie)
	}
	c.SetCookie("session_id", "", -1, "/", "", false, true)
	c.String(200, "Logged out successfully.")
}


func main() {
	r := gin.Default()

	r.Use(LoggerMiddleware())

	r.GET("/login", Login)

	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	auth.GET("/dashboard", Dashboard)
	auth.GET("/logout", Logout)

	r.Run(":8080")
}

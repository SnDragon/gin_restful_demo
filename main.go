package main

import (
	"github.com/SnDragon/gin_restful_demo/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "pong"})
	})
	r.POST("/users", api.CreateUser)
	r.PUT("/users/:id", api.UpdateUser)
	r.DELETE("/users/:id", api.DeleteUser)
	r.GET("/users/:id", api.GetUser)
	r.GET("/users", api.GetUsers)
	r.Run()
}

package main

import (
	"github.com/adetiamarhadi/golang-sqlx/controller"
	"github.com/adetiamarhadi/golang-sqlx/dbclient"
	"github.com/gin-gonic/gin"
)

func main() {
	dbclient.InitialiseDBConnection()

	r := gin.Default()

	r.POST("/", controller.CreatePost)
	r.GET("/", controller.GetPosts)
	r.GET("/:id", controller.GetPost)

	if err := r.Run(":5000"); err != nil {
		panic(err.Error())
	}
}
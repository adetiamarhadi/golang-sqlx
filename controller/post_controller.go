package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adetiamarhadi/golang-sqlx/dbclient"
	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int64 `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func CreatePost(c *gin.Context) {
	var reqBody Post
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": true,
			"message": "Invalid request body",
		})
		return
	}

	res, _ := dbclient.DBClient.Exec("INSERT INTO posts (title, content) VALUES (?, ?);", reqBody.Title, reqBody.Content)

	id, _ := res.LastInsertId()

	c.JSON(http.StatusCreated, gin.H{
		"error": false,
		"id": id,
	})
}

func GetPosts(c *gin.Context) {
	var posts []Post

	dbclient.DBClient.Select(&posts, "SELECT id, title, content, created_at FROM posts")

	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var post Post

	dbclient.DBClient.Get(&post, "SELECT id, title, content, created_at FROM posts WHERE id = ?", id)

	c.JSON(http.StatusOK, post)
}
package controller

import (
	"database/sql"
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
	CreatedAt time.Time `json:"created_at"`
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

	rows, err := dbclient.DBClient.Query("SELECT id, title, content, created_at FROM posts")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": true,
		})
		return
	}

	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": true,
			})
			return
		}

		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	row := dbclient.DBClient.QueryRow("SELECT id, title, content, created_at FROM posts WHERE id = ?", id)

	var post Post
	if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": true,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
		return
	}

	c.JSON(http.StatusOK, post)
}
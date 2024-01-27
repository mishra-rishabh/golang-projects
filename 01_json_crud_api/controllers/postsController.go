package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mishra-rishabh/json_crud_api/initializers"
	"github.com/mishra-rishabh/json_crud_api/models"
)

func PostsCreate(c *gin.Context) {
	/* Get data of request body */
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	/* Create a post */
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	/* Return it */
	c.JSON(200, gin.H{
		"post": post,
	})
}

func FetchAllPosts(c *gin.Context) {
	/* Get all posts */
	var posts []models.Post
	initializers.DB.Find(&posts)

	/* Respond with all the posts */
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func FetchPost(c *gin.Context) {
	/* Get post id from url */
	id := c.Param("id")

	/* Get the post */
	var post models.Post
	initializers.DB.First(&post, id)

	/* Respond with the post */
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	/* Get post id from url */
	id := c.Param("id")

	/* Get the data of request body */
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	/* Find the post we have to update */
	var post models.Post
	initializers.DB.First(&post, id)

	/* Update the post */
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	/* Respond with the post */
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	/* Get post id from url */
	id := c.Param("id")

	/* Delete the post */
	initializers.DB.Delete(&models.Post{}, id)

	/* Respond */
	c.JSON(200, gin.H{
		"message": "Post deleted successfully!",
	})
}

package controllers

import (
	"github.com/alpish/go-crud/intializers"
	"github.com/alpish/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	ctx.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := intializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	// Return post

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(ctx *gin.Context) {
	// Get all records
	var posts []models.Post
	intializers.DB.Find(&posts)

	// Respond with them
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(ctx *gin.Context) {
	// Get id of URL
	id := ctx.Param("id")

	// Get record
	var post models.Post
	intializers.DB.First(&post, id)

	// Respond with them
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(ctx *gin.Context) {
	// Get id of URL
	id := ctx.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	ctx.Bind(&body)

	// Find post
	var post models.Post
	intializers.DB.First(&post, id)

	intializers.DB.Model(&post).Updates(models.Post{Title: "hello", Body: "34"})

	// Respond with them
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(ctx *gin.Context) {
	// Get id of URL
	id := ctx.Param("id")

	// Delete posts
	intializers.DB.Delete(&models.Post{}, id)

	// Respond
	ctx.Status(200)
}

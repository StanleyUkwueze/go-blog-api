package controllers

import (
	loadinitializers "apicrud/LoadInitializers"
	"apicrud/models"
	"log"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := loadinitializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostFetch(c *gin.Context) {
	var posts []models.Post
	loadinitializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostFetchById(c *gin.Context) {

	id := c.Param("id")
	var post models.Post

	loadinitializers.DB.Find(&post, id)

	c.JSON(200, gin.H{
		"post/": post,
	})
}

func PostUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post

	loadinitializers.DB.Find(&post, id)

	loadinitializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	result := loadinitializers.DB.Find(&models.Post{}, id)

	if result.RowsAffected == 0 {
		log.Fatal("No record found")
		return
	}
	loadinitializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		log.Fatal("An error occured")
	}
	c.Status(200)
}

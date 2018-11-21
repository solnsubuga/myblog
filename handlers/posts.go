package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsubuga/myblog/models"
)

//PostHandler handles requests for posts
type PostHandler struct {
	Repo *models.PostRepository
}

//HandlePosts handles get all posts
func (h *PostHandler) HandlePosts(c *gin.Context) {
	posts := h.Repo.GetPosts()
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

//CreatePost handles creating a post
func (h *PostHandler) CreatePost(c *gin.Context) {
	post := &models.Post{}
	c.BindJSON(post)
	//TODO validate data
	if err := h.Repo.Save(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

//GetPost gets a specific post by id
func (h *PostHandler) GetPost(c *gin.Context) {
	id := c.Param("id")
	post, err := h.Repo.GetPost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

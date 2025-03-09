package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakihaha/go-forum/internal/model/posts"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var req posts.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	ctx := c.Request.Context()

	err := h.postSvc.CreatePost(ctx, userID.(int64), req)
	// if err := h.postSvc.CreatePost(ctx, userID.(int64), req); err != nil {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created"})
}

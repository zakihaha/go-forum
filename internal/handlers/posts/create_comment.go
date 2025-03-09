package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zakihaha/go-forum/internal/model/posts"
)

func (h *Handler) CreateComment(c *gin.Context) {
	var req posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid postID").Error()})
		return
	}

	userID := c.GetInt64("userID")
	ctx := c.Request.Context()

	err = h.postSvc.CreateComment(ctx, postID, userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment created"})
}

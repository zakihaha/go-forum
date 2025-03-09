package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPostByID(c *gin.Context) {
	id := c.Param("postID")

	postID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid postID",
		})
		return
	}

	post, err := h.postSvc.GetPostByID(c.Request.Context(), int64(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get post by ID",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

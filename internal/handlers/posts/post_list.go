package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()

	pageSize := c.DefaultQuery("pageSize", "10")
	pageIndex := c.DefaultQuery("pageIndex", "1")

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid pageSize",
		})
		return
	}

	pageIndexInt, err := strconv.Atoi(pageIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid pageIndex",
		})
		return
	}

	posts, err := h.postSvc.GetAllPost(ctx, pageSizeInt, pageIndexInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get all post",
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

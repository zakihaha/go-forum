package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakihaha/go-forum/internal/model/memberships"
)

func (h *Handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	token, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := memberships.RefreshResponse{
		AccessToken: token,
	}

	c.JSON(http.StatusOK, response)
}

package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/zakihaha/go-forum/internal/configs"
	"github.com/zakihaha/go-forum/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.GetConfig().Service.SecretJWT
	return func(c *gin.Context) {
		// Get the token from the header
		header := c.GetHeader("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		// Check if the token is valid
		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)

		c.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.GetConfig().Service.SecretJWT
	return func(c *gin.Context) {
		// Get the token from the header
		header := c.GetHeader("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		// Check if the token is valid
		userID, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
		log.Info().Msgf("userID: %d, username: %s", userID, username)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)

		c.Next()
	}
}

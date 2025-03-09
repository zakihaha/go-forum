package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/zakihaha/go-forum/internal/model/memberships"
	"github.com/zakihaha/go-forum/pkg/jwt"
	tokenUtil "github.com/zakihaha/go-forum/pkg/token"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
		return "", "", nil
	}

	if user == nil {
		return "", "", errors.New("email nor exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", "", errors.New("password not match")
	}

	// secretKey := configs.GetConfig().Service.SecretJWT ??

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return "", "", nil
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token")
		return "", "", nil
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	refreshToken := tokenUtil.GenerateRefreshToken()
	if refreshToken == "" {
		log.Error().Msg("Failed to generate refresh token")
		return "", "", nil
	}

	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(time.Hour * 24 * 7),
		CreatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(user.ID, 10),
		UpdatedAt:    time.Now(),
		UpdatedBy:    strconv.FormatInt(user.ID, 10),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert refresh token")
		return "", "", nil
	}

	return token, refreshToken, nil
}

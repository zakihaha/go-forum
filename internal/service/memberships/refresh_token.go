package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zakihaha/go-forum/internal/model/memberships"
	"github.com/zakihaha/go-forum/pkg/jwt"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	refreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token")
		return "", err
	}

	if refreshToken == nil {
		return "", errors.New("refresh token expired aaa")
	}

	// if refreshToken.ExpiredAt.Before(request.Now) {
	// 	return "", memberships.ErrRefreshTokenExpired
	// }

	if refreshToken.RefreshToken != request.Token {
		return "", errors.New("refresh token not match")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
		return "", nil
	}

	if user == nil {
		return "", errors.New("user not exist")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return "", err
	}

	return token, nil
}

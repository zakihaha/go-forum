package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zakihaha/go-forum/internal/model/posts"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedAt: now,
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user activity")
		return err
	}

	if userActivity == nil {
		// create new user activity
		if !request.IsLiked {
			return errors.New("please like the post first")
		}

		err = s.postRepo.CreateUserActivity(ctx, model)
	} else {
		// update user activity
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("Failed to upsert user activity")
		return err
	}

	return nil
}

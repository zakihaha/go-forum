package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zakihaha/go-forum/internal/model/posts"
)

func (s *service) CreateComment(ctx context.Context, postID, userID int64, request posts.CreateCommentRequest) error {
	now := time.Now()

	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: request.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}

	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create comment")
		return err
	}

	return nil
}

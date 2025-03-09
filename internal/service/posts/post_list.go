package posts

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zakihaha/go-forum/internal/model/posts"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)

	response, err := s.postRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all post")
		return response, err
	}

	return response, nil
}

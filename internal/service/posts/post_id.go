package posts

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zakihaha/go-forum/internal/model/posts"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get post by ID a")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to count like by post ID")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get comment by post ID")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil

}

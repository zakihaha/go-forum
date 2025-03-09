package posts

import (
	"context"

	"github.com/zakihaha/go-forum/internal/configs"
	"github.com/zakihaha/go-forum/internal/model/posts"
)

type postRepository interface {
	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error)
	GetPostByID(ctx context.Context, id int64) (*posts.Post, error)
	CreatePost(ctx context.Context, model posts.PostModel) error

	CreateComment(ctx context.Context, model posts.CommentModel) error
	GetCommentByPostID(ctx context.Context, postID int64) ([]posts.Comment, error)

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	CountLikeByPostID(ctx context.Context, postID int64) (int, error)
}

type service struct {
	cfg      *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}

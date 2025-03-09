package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/zakihaha/go-forum/internal/middleware"
	"github.com/zakihaha/go-forum/internal/model/posts"
)

type postService interface {
	GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error)
	GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error)

	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userID int64, request posts.CreateCommentRequest) error

	UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error
}

type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")

	route.Use(middleware.AuthMiddleware())
	route.GET("/", h.GetAllPost)
	route.POST("/create", h.CreatePost)
	route.GET("/:postID", h.GetPostByID)
	route.POST("/:postID/comment", h.CreateComment)
	route.PUT("/:postID/user-activity", h.UpsertUserActivity)
}

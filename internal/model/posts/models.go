package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle    string   `json:"postTitle"`
		PostContent  string   `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
	}
)

type (
	PostModel struct {
		ID           int64     `db:"id"`
		UserID       int64     `db:"user_id"`
		PostTitle    string    `db:"post_title"`
		PostContent  string    `db:"post_content"`
		PostHashtags string    `db:"post_hashtags"`
		CreatedAt    time.Time `db:"created_at"`
		CreatedBy    string    `db:"created_by"`
		UpdatedAt    time.Time `db:"updated_at"`
		UpdatedBy    string    `db:"updated_by"`
	}
)

type (
	GetAllPostResponse struct {
		Data       []Post     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	GetPostResponse struct {
		PostDetail Post      `json:"dataDetail"`
		LikeCount  int       `json:"likeCount"`
		Comments   []Comment `json:"comments"`
	}

	Post struct {
		ID           int64    `json:"id"`
		UserID       int64    `json:"UserID"`
		Username     string   `json:"username"`
		PostTitle    string   `json:"postTitle"`
		PostContent  string   `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
		IsLiked      bool     `json:"isLiked"`
	}

	Comment struct {
		ID             int64  `json:"id"`
		UserID         int64  `json:"userID"`
		Username       string `json:"username"`
		CommentContent string `json:"commentContent"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
)

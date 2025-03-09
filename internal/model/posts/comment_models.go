package posts

import "time"

type (
	CreateCommentRequest struct {
		CommentContent string `json:"comment_content"`
	}
)

type (
	CommentModel struct {
		ID             int64     `db:"id"`
		PostID         int64     `db:"post_id"`
		UserID         int64     `db:"user_id"`
		CommentContent string    `db:"comment_content"`
		CreatedAt      time.Time `db:"created_at"`
		CreatedBy      string    `db:"created_by"`
		UpdatedAt      time.Time `db:"updated_at"`
		UpdatedBy      string    `db:"updated_by"`
	}
)

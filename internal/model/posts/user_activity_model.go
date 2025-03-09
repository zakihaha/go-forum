package posts

import "time"

type (
	UserActivityRequest struct {
		IsLiked bool `json:"isLiked"`
	}
)

type (
	UserActivityModel struct {
		ID        int64     `db:"id"`
		PostID    int64     `db:"post_id"`
		UserID    int64     `db:"user_id"`
		IsLiked   bool      `db:"is_liked"`
		CreatedAt time.Time `db:"created_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedAt time.Time `db:"updated_at"`
		UpdatedBy string    `db:"updated_by"`
	}
)

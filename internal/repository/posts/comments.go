package posts

import (
	"context"

	"github.com/zakihaha/go-forum/internal/model/posts"
)

func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments (post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetCommentByPostID(ctx context.Context, postID int64) ([]posts.Comment, error) {
	query := `SELECT c.id, c.user_id, u.username, c.comment_content
				FROM comments c 
				JOIN users u ON c.user_id = u.id
				WHERE c.post_id = ?
				ORDER BY c.updated_at DESC`
	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []posts.Comment
	for rows.Next() {
		var (
			model    posts.Comment
			username string
		)

		err := rows.Scan(&model.ID, &model.UserID, &username, &model.CommentContent)
		if err != nil {
			return nil, err
		}

		data = append(data, posts.Comment{
			ID:             model.ID,
			UserID:         model.UserID,
			Username:       username,
			CommentContent: model.CommentContent,
		})
	}

	return data, nil
}

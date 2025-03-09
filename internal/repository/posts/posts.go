package posts

import (
	"context"
	"strings"

	"github.com/zakihaha/go-forum/internal/model/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO posts (user_id, post_title, post_content, post_hashtags, created_at, created_by, updated_at, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostTitle, model.PostContent, model.PostHashtags, model.CreatedAt, model.CreatedBy, model.UpdatedAt, model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error) {
	response := posts.GetAllPostResponse{}

	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hashtags
				FROM posts p JOIN users u ON p.user_id = u.id
				ORDER BY p.updated_at DESC
				LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	var data []posts.Post
	for rows.Next() {
		var (
			model    posts.PostModel
			username string
		)

		err := rows.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHashtags)
		if err != nil {
			return response, err
		}

		data = append(data, posts.Post{
			ID:           model.ID,
			UserID:       model.UserID,
			Username:     username,
			PostTitle:    model.PostTitle,
			PostContent:  model.PostContent,
			PostHashtags: strings.Split(model.PostHashtags, ","),
		})
	}

	response.Data = data
	response.Pagination = posts.Pagination{
		Limit:  limit,
		Offset: offset,
	}

	return response, nil
}

func (r *repository) GetPostByID(ctx context.Context, id int64) (*posts.Post, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hashtags, IFNULL(ua.is_liked, 0)
				FROM posts p
				JOIN users u ON p.user_id = u.id
				LEFT JOIN user_activities ua ON p.id = ua.post_id
				WHERE p.id = ?`

	row := r.db.QueryRowContext(ctx, query, id)

	var (
		model    posts.PostModel
		username string
		isLiked  bool
	)
	err := row.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHashtags, &isLiked)
	if err != nil {
		return &posts.Post{}, err
	}

	response := &posts.Post{
		ID:           model.ID,
		UserID:       model.UserID,
		Username:     username,
		PostTitle:    model.PostTitle,
		PostContent:  model.PostContent,
		PostHashtags: strings.Split(model.PostHashtags, ","),
		IsLiked:      isLiked,
	}

	return response, nil
}

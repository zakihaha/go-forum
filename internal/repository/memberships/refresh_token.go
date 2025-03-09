package memberships

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zakihaha/go-forum/internal/model/memberships"
)

func (r *repository) InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, expired_at, created_at, created_by, updated_at, updated_by)
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.ExpiredAt, model.CreatedAt, model.CreatedBy, model.UpdatedAt, model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error) {
	log.Info().Msg("GetRefreshToken" + strconv.FormatInt(userID, 10))
	var response memberships.RefreshTokenModel
	query := `SELECT id, user_id, refresh_token, expired_at, created_at, created_by, updated_at, updated_by
		FROM refresh_tokens 
		WHERE user_id = ? 
		AND expired_at >= ?`

	row := r.db.QueryRowContext(ctx, query, userID, now)
	err := row.Scan(&response.ID, &response.UserID, &response.RefreshToken, &response.ExpiredAt, &response.CreatedAt, &response.CreatedBy, &response.UpdatedAt, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}

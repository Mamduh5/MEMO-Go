package mysql

import (
	"context"
	"database/sql"
	"errors"

	"memo-go/services/auth/internal/domain"
)

type RefreshTokenRepository struct {
	db *sql.DB
}

func NewRefreshTokenRepository(db *sql.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Save(
	ctx context.Context,
	token *domain.RefreshToken,
) error {

	const query = `
INSERT INTO refresh_tokens (
    id,
    user_id,
    token,
    expires_at,
    revoked
) VALUES (?, ?, ?, ?, ?)
`

	_, err := r.db.ExecContext(
		ctx,
		query,
		token.ID,
		token.UserID,
		token.Token,
		token.ExpiresAt,
		token.Revoked,
	)

	return err
}

func (r *RefreshTokenRepository) Find(
	ctx context.Context,
	token string,
) (*domain.RefreshToken, error) {

	const query = `
SELECT id, user_id, token, expires_at, revoked
FROM refresh_tokens
WHERE token = ?
LIMIT 1
`

	row := r.db.QueryRowContext(ctx, query, token)

	rt := &domain.RefreshToken{}
	if err := row.Scan(
		&rt.ID,
		&rt.UserID,
		&rt.Token,
		&rt.ExpiresAt,
		&rt.Revoked,
	); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return rt, nil
}

func (r *RefreshTokenRepository) Revoke(
	ctx context.Context,
	tokenID string,
) error {

	const query = `
UPDATE refresh_tokens
SET revoked = true
WHERE id = ?
`

	_, err := r.db.ExecContext(ctx, query, tokenID)
	return err
}

func (r *RefreshTokenRepository) RevokeAllByUser(
	ctx context.Context,
	userID string,
) error {

	const query = `
UPDATE refresh_tokens
SET revoked = true
WHERE user_id = ? AND revoked = false
`

	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

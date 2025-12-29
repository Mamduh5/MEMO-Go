package mysql

import (
	"context"
	"database/sql"
	"errors"

	"memo-go/services/auth/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {

	const query = `
SELECT id, email, password_hash
FROM users
WHERE email = ?
LIMIT 1
`

	row := r.db.QueryRowContext(ctx, query, email)

	u := &domain.User{}
	if err := row.Scan(&u.ID, &u.Email, &u.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {

	const query = `
INSERT INTO users (id, email, password_hash)
VALUES (?, ?, ?)
`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.ID,
		user.Email,
		user.Password,
	)

	return err
}

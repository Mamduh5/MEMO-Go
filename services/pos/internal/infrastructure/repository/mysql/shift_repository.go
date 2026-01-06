package mysql

import (
	"context"
	"database/sql"
	"time"

	"memo-go/services/pos/internal/domain"
)

type ShiftRepository struct {
	db *sql.DB
}

func NewShiftRepository(db *sql.DB) *ShiftRepository {
	return &ShiftRepository{db: db}
}

func (r *ShiftRepository) Create(ctx context.Context, s *domain.Shift) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO shifts (id, user_id, opened_at, closed_at)
		 VALUES (?, ?, ?, ?)`,
		s.ID,
		s.UserID,
		s.OpenedAt,
		s.ClosedAt,
	)
	return err
}

func (r *ShiftRepository) FindOpenByUserID(
	ctx context.Context,
	userID string,
) (*domain.Shift, error) {

	row := r.db.QueryRowContext(
		ctx,
		`SELECT id, user_id, opened_at, closed_at
		 FROM shifts
		 WHERE user_id = ? AND closed_at IS NULL
		 LIMIT 1`,
		userID,
	)

	var s domain.Shift
	var closedAt sql.NullTime

	err := row.Scan(
		&s.ID,
		&s.UserID,
		&s.OpenedAt,
		&closedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if closedAt.Valid {
		s.ClosedAt = &closedAt.Time
	}

	return &s, nil
}

func (r *ShiftRepository) Close(
	ctx context.Context,
	shiftID string,
	closedAt time.Time,
) error {

	res, err := r.db.ExecContext(
		ctx,
		`UPDATE shifts
		 SET closed_at = ?
		 WHERE id = ? AND closed_at IS NULL`,
		closedAt,
		shiftID,
	)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

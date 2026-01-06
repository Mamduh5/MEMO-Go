package pos

import (
	"context"
	"memo-go/services/pos/internal/domain"
	"time"

	"github.com/google/uuid"
)

type OpenShiftResult struct {
	ShiftID  string
	OpenedAt time.Time
}

func (u *PosUsecase) OpenShift(
	ctx context.Context,
	userID string,
) (*OpenShiftResult, error) {

	// rule (will be enforced strictly next step)
	existing, err := u.shiftRepo.FindOpenByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrShiftAlreadyOpen
	}

	now := time.Now().UTC()

	shift := &domain.Shift{
		ID:       uuid.NewString(),
		UserID:   userID,
		OpenedAt: now,
	}

	if err := u.shiftRepo.Create(ctx, shift); err != nil {
		return nil, err
	}

	return &OpenShiftResult{
		ShiftID:  shift.ID,
		OpenedAt: shift.OpenedAt,
	}, nil
}

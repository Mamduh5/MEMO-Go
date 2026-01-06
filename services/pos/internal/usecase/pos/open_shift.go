package pos

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type OpenShiftResult struct {
	ShiftID  string
	OpenedAt time.Time
}

func (u *Usecase) OpenShift(
	ctx context.Context,
	userID string,
) (*OpenShiftResult, error) {

	// Business rule placeholder:
	// - one open shift per user
	// - persist later

	now := time.Now().UTC()

	return &OpenShiftResult{
		ShiftID:  uuid.NewString(),
		OpenedAt: now,
	}, nil
}

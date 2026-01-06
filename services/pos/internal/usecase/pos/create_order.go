package pos

import (
	"context"
	"time"

	"memo-go/services/pos/internal/domain"

	"github.com/google/uuid"
)

func (u *PosUsecase) CreateOrder(
	ctx context.Context,
	userID string,
) (string, error) {

	shift, err := u.shiftRepo.FindOpenByUserID(ctx, userID)
	if err != nil {
		return "", err
	}
	if shift == nil {
		return "", ErrNoOpenShift
	}

	order := &domain.Order{
		ID:        uuid.NewString(),
		UserID:    userID,
		ShiftID:   shift.ID,
		Status:    domain.OrderStatusOpen,
		CreatedAt: time.Now().UTC(),
	}

	if err := u.orderRepo.Create(ctx, order); err != nil {
		return "", err
	}

	return order.ID, nil
}

package pos

import (
	"context"
	"database/sql"
	"memo-go/services/pos/internal/domain"
	"time"
)

func (u *PosUsecase) CloseOrder(
	ctx context.Context,
	orderID string,
) (int64, error) {

	order, err := u.orderRepo.FindByID(ctx, orderID)
	if err != nil {
		return 0, err
	}
	if order == nil {
		return 0, ErrOrderNotFound
	}
	if order.Status != domain.OrderStatusOpen {
		return 0, ErrOrderClosed
	}

	items, err := u.itemRepo.ListByOrderID(ctx, orderID)
	if err != nil {
		return 0, err
	}
	if len(items) == 0 {
		return 0, ErrEmptyOrder
	}

	var total int64
	for _, it := range items {
		total += it.Price * int64(it.Quantity)
	}

	now := time.Now().UTC()

	if err := u.orderRepo.Close(ctx, orderID, total, now); err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrOrderClosed
		}
		return 0, err
	}

	return total, nil
}

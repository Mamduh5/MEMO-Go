package pos

import (
	"context"

	"memo-go/services/pos/internal/domain"

	"github.com/google/uuid"
)

func (u *PosUsecase) AddOrderItem(
	ctx context.Context,
	orderID string,
	name string,
	price int64,
	quantity int,
) error {

	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	if price < 0 {
		return ErrInvalidPrice
	}

	order, err := u.orderRepo.FindByID(ctx, orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return ErrOrderNotFound
	}
	if order.Status != domain.OrderStatusOpen {
		return ErrOrderClosed
	}

	item := &domain.OrderItem{
		ID:       uuid.NewString(),
		OrderID:  orderID,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}

	return u.itemRepo.Add(ctx, item)
}

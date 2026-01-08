package pos

import "memo-go/services/pos/internal/domain"

type PosUsecase struct {
	shiftRepo domain.ShiftRepository
	orderRepo domain.OrderRepository
	itemRepo  domain.OrderItemRepository
}

func NewPosUsecase(
	shiftRepo domain.ShiftRepository,
	orderRepo domain.OrderRepository,
	itemRepo domain.OrderItemRepository,

) *PosUsecase {
	return &PosUsecase{
		shiftRepo: shiftRepo,
		orderRepo: orderRepo,
		itemRepo:  itemRepo,
	}
}

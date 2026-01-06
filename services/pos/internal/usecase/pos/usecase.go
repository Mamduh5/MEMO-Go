package pos

import "memo-go/services/pos/internal/domain"

type PosUsecase struct {
	shiftRepo domain.ShiftRepository
	orderRepo domain.OrderRepository
	itemRepo  domain.OrderItemRepository
}

func NewPosUsecase(shiftRepo domain.ShiftRepository) *PosUsecase {
	return &PosUsecase{
		shiftRepo: shiftRepo,
	}
}

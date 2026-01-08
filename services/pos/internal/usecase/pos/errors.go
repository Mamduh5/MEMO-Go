package pos

import "errors"

var ErrShiftAlreadyOpen = errors.New("shift already open")
var ErrNoOpenShift = errors.New("no open shift")
var (
	ErrOrderNotFound   = errors.New("order not found")
	ErrOrderClosed     = errors.New("order is closed")
	ErrInvalidQuantity = errors.New("invalid quantity")
	ErrInvalidPrice    = errors.New("invalid price")
)

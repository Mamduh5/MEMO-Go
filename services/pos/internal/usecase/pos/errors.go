package pos

import "errors"

var ErrShiftAlreadyOpen = errors.New("shift already open")
var ErrNoOpenShift = errors.New("no open shift")

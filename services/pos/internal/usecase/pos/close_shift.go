package pos

import (
	"context"
	"database/sql"
	"time"
)

func (u *PosUsecase) CloseShift(
	ctx context.Context,
	userID string,
) error {

	shift, err := u.shiftRepo.FindOpenByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if shift == nil {
		return ErrNoOpenShift
	}

	now := time.Now().UTC()

	if err := u.shiftRepo.Close(ctx, shift.ID, now); err != nil {
		if err == sql.ErrNoRows {
			return ErrNoOpenShift
		}
		return err
	}

	return nil
}

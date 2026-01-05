package app

import (
	"memo-go/services/pos/internal/config"
)

func Run(cfg *config.Config) error {
	return startGRPCServer(cfg)
}

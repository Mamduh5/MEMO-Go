package app

import (
	"memo-go/services/pos/internal/config"
	dbMysql "memo-go/services/pos/internal/infrastructure/db/mysql"
	repoMysql "memo-go/services/pos/internal/infrastructure/repository/mysql"
	"memo-go/services/pos/internal/usecase/pos"
)

func NewApp() error {
	cfg := config.Load()

	db, err := dbMysql.New(dbMysql.Config{
		User:     cfg.MySQL.User,
		Password: cfg.MySQL.Password,
		Host:     cfg.MySQL.Host,
		Port:     cfg.MySQL.Port,
		DBName:   cfg.MySQL.DBName,
	})
	if err != nil {
		return err
	}

	if err := dbMysql.Migrate(db); err != nil {
		return err
	}

	shiftRepo := repoMysql.NewShiftRepository(db)
	orderRepo := repoMysql.NewOrderRepository(db)
	itemRepo := repoMysql.NewOrderItemRepository(db)
	posUC := pos.NewPosUsecase(
		shiftRepo,
		orderRepo,
		itemRepo,
	)

	return StartGRPCServer(cfg, posUC)
}

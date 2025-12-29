package app

import (
	"context"
	"fmt"

	"memo-go/services/auth/internal/config"
	dbMysql "memo-go/services/auth/internal/infrastructure/db/mysql"
	repoMysql "memo-go/services/auth/internal/infrastructure/repository/mysql"
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
	userRepo := repoMysql.NewUserRepository(db)

	ctx := context.Background()
	user, err := userRepo.FindByEmail(ctx, "test@example.com")
	fmt.Println(user, err)

	return nil
}

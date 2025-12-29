package app

import (
	"memo-go/services/auth/internal/config"
	dbMysql "memo-go/services/auth/internal/infrastructure/db/mysql"
	repoMysql "memo-go/services/auth/internal/infrastructure/repository/mysql"
	bcryptHasher "memo-go/services/auth/internal/infrastructure/security/bcrypt"
	"memo-go/services/auth/internal/usecase/auth"
	"time"

	jwtToken "memo-go/services/auth/internal/infrastructure/token/jwt"
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
	tokenRepo := repoMysql.NewRefreshTokenRepository(db)
	hasher := bcryptHasher.New(0)

	tokenGen := jwtToken.New(
		"dev-secret-change-later",
		15*time.Minute,
	)

	authUC := auth.NewAuthUsecase(
		userRepo,
		tokenRepo,
		hasher,
		tokenGen,
	)

	return StartGRPCServer(authUC)
}

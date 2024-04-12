package app

import (
	"github.com/horasachy/auth/internal/usecase"
	"github.com/horasachy/auth/pkg/config"
	"github.com/horasachy/auth/pkg/storage/mongodb"
)

func Run(cfg *config.Config) {
	mongoDb, _ := mongodb.New(cfg.MongoUrl)
	ucc := &usecase.UserUseCase{userRepo: mongoClient}
}

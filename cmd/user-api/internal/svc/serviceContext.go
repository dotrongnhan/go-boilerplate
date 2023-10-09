package svc

import (
	"git.savvycom.vn/go-services/go-boilerplate/cmd/user-api/internal/config"
	"git.savvycom.vn/go-services/go-boilerplate/internal/app/usecases"
	"git.savvycom.vn/go-services/go-boilerplate/internal/infrastructure/database"
	"github.com/zeromicro/go-zero/core/service"
)

type ServiceContext struct {
	Config      config.Config
	UserUseCase usecases.UserUseCase
}

func NewServiceContext(c config.Config) *ServiceContext {

	userRepo := database.NewUserRepository(c.DataSource, c.Mode == service.DevMode)
	userUseCase := usecases.NewUserUsecase(userRepo)

	return &ServiceContext{
		Config:      c,
		UserUseCase: userUseCase,
	}
}

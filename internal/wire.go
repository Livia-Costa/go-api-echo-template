//go:build wireinject
// +build wireinject

//go:generate wire

package internal

import (
	"github.com/google/wire"
	"github.com/victorsantosbrazil/go-api-echo-template/internal/app/api/rest"
	"github.com/victorsantosbrazil/go-api-echo-template/internal/app/usecase"
	"github.com/victorsantosbrazil/go-api-echo-template/internal/config"
	"github.com/victorsantosbrazil/go-api-echo-template/internal/domain/service"
	"github.com/victorsantosbrazil/go-api-echo-template/internal/infra/dao"
	"github.com/victorsantosbrazil/go-api-echo-template/pkg/app/echo"
	"github.com/victorsantosbrazil/go-api-echo-template/pkg/infra/datasource/postgresql"
)

var (
	DAOSet        = wire.NewSet(dao.NewInstitutionDAO)
	ServiceSet    = wire.NewSet(service.NewInstitutionService)
	UseCaseSet    = wire.NewSet(usecase.NewListInstitutionsUseCase)
	ControllerSet = wire.NewSet(rest.NewV1Group, rest.NewInstitutionController)
)

func NewApp() (*App, error) {
	wire.Build(
		config.ReadConfig,
		GetDataSourceConfig,
		echo.New,
		DAOSet,
		ServiceSet,
		UseCaseSet,
		ControllerSet,
		rest.NewControllers,
		newApp,
	)

	return &App{}, nil
}

func GetDataSourceConfig(cfg *config.Config) *postgresql.Config {
	return cfg.DataSource
}

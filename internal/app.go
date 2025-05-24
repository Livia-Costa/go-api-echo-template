package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/go-api-echo-template/internal/app/api/rest"
	"github.com/victorsantosbrazil/go-api-echo-template/internal/config"
	"github.com/victorsantosbrazil/go-api-echo-template/pkg/app/log"
	"github.com/victorsantosbrazil/go-api-echo-template/pkg/infra/datasource/postgresql"
)

// App represents the main application structure
type App struct {
	echo        *echo.Echo
	logger      log.Logger
	config      *config.Config
	controllers *rest.Controllers
}

func (a *App) Start() error {
	err := a.setupDatabase()
	if err != nil {
		return err
	}

	addr := a.config.GetAddress()
	return a.echo.Start(addr)
}

func (a *App) setupDatabase() error {
	dsConfig := a.config.DataSource

	m, err := postgresql.NewMigration(dsConfig)
	if err != nil {
		return err
	}

	m.Up()
	return nil
}

func newApp(e *echo.Echo, cfg *config.Config, controllers *rest.Controllers) *App {
	return &App{echo: e, config: cfg, controllers: controllers}
}

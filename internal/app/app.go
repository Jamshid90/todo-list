package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	deliveryHttp "github.com/Jamshid90/todo-list/internal/delivery/http"
	"github.com/Jamshid90/todo-list/internal/infrastructure/repository"
	"github.com/Jamshid90/todo-list/internal/pkg/config"
	"github.com/Jamshid90/todo-list/internal/pkg/postgres"
	"github.com/Jamshid90/todo-list/internal/usecase"
	"go.uber.org/zap"
)

const (
	EnvironmentProduction = "production"
	EnvironmentDevelop    = "develop"
)

type App struct {
	config *config.Config
	logger *zap.Logger
	db     *postgres.PostgresDB
	server *http.Server
}

func NewApp(cfg *config.Config, logger *zap.Logger) (*App, error) {
	// initialization db
	db, err := postgres.New(
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Sslmode,
	)
	if err != nil {
		return nil, err
	}

	return &App{
		config: cfg,
		logger: logger,
		db:     db,
	}, nil
}

func (a *App) Run() error {
	// context timeout initialization
	contextTimeout, err := time.ParseDuration(a.config.Context.Timeout)
	if err != nil {
		return fmt.Errorf("error during parse duration for context timeout : %w", err)
	}

	taskWrite := repository.NewTaskWrite(a.db)
	taskRead := repository.NewTaskRead(a.db)

	taskUseCase := usecase.NewTask(taskWrite, taskRead)

	// initialize api handler
	apiHandler := deliveryHttp.NewRoute(a.config, a.logger, contextTimeout, taskUseCase)

	// initialize api server
	a.server, err = deliveryHttp.NewServer(a.config, apiHandler)
	if err != nil {
		return fmt.Errorf("error during http server initialization: %w", err)
	}

	address := fmt.Sprintf("%s:%s", a.config.HttpServer.Host, a.config.HttpServer.Port)
	swaggerAddress := fmt.Sprintf("%s:%s/%s", a.config.HttpServer.Host, a.config.HttpServer.Port, "swagger/index.html")

	a.logger.Info("listen: ", zap.String("address", address))
	a.logger.Info("swagger: ", zap.String("address", swaggerAddress))
	return a.server.ListenAndServe()
}

func (a *App) Stop() error {
	a.db.Close()
	// shutdown server http
	if err := a.server.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("error when shutting down http server: %w", err)
	}
	return nil
}

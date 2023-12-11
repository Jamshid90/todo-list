package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	apppkg "github.com/Jamshid90/todo-list/internal/app"
	"github.com/Jamshid90/todo-list/internal/pkg/config"
	loggerpkg "github.com/Jamshid90/todo-list/internal/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use: "task-server",
	Run: func(cmd *cobra.Command, args []string) {
		// initialization config
		cfg, err := config.New()
		if err != nil {
			log.Fatal(err)
		}
		//initialization logger
		logger, err := loggerpkg.NewDevelopment(cfg.LogLevel, cfg.App)
		if apppkg.EnvironmentProduction == cfg.Environment {
			logger, err = loggerpkg.NewProduction(cfg.LogLevel, cfg.App)
		}
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err := logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
				log.Panicf("error during sync logger: %s", err.Error())
			}
		}()
		//initialization application
		app, err := apppkg.NewApp(cfg, logger)
		if err != nil {
			log.Fatal(err)
		}

		stop := make(chan os.Signal, 1)
		go func() {
			if err := app.Run(); err != nil {
				logger.Error("error during application launch", zap.Error(err))
			}
		}()

		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		sign := <-stop

		logger.Info("stopping application", zap.String("signal", sign.String()))
		if err := app.Stop(); err != nil {
			logger.Error("error during stop application", zap.Error(err))
		}
		logger.Info("application stopped")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

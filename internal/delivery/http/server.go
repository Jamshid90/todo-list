package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Jamshid90/todo-list/internal/pkg/config"
)

func NewServer(config *config.Config, handler http.Handler) (*http.Server, error) {
	// read timeout initialization
	readTimeout, err := time.ParseDuration(config.HttpServer.ReadTimeout)
	if err != nil {
		return nil, fmt.Errorf("error during parse duration for server read timeout: %w", err)
	}
	// write timeout initialization
	writeTimeout, err := time.ParseDuration(config.HttpServer.WriteTimeout)
	if err != nil {
		return nil, fmt.Errorf("error during parse duration for server write timeout: %w", err)
	}
	// idle timeout initialization
	idleTimeout, err := time.ParseDuration(config.HttpServer.IdleTimeout)
	if err != nil {
		return nil, fmt.Errorf("error during parse duration for server idle timeout: %w", err)
	}

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", config.HttpServer.Host, config.HttpServer.Port),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      handler,
	}, nil
}

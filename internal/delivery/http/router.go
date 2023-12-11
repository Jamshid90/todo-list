package api

import (
	"net/http"
	"time"

	_ "github.com/Jamshid90/todo-list/internal/delivery/http/docs"
	handlerv1 "github.com/Jamshid90/todo-list/internal/delivery/http/handler/v1"
	middlewarepkg "github.com/Jamshid90/todo-list/internal/delivery/http/middleware"
	"github.com/Jamshid90/todo-list/internal/pkg/config"
	"github.com/Jamshid90/todo-list/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
)

// @title To Do List API
// @version 1.0
// @BasePath /v1
func NewRoute(cfg *config.Config, logger *zap.Logger, contextTimeout time.Duration, taskUseCase usecase.TaskService) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(contextTimeout))
	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Request-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Route("/v1", func(r chi.Router) {
		r.Use(middlewarepkg.ContentTypeJson)
		r.Mount("/task", handlerv1.NewTaskHandler(logger, taskUseCase))
	})

	// declare swagger api route
	r.Get("/swagger/*", httpSwagger.Handler())
	return r
}

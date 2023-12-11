package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Jamshid90/todo-list/internal/delivery/http/models"
	"github.com/Jamshid90/todo-list/internal/entity"
	"github.com/Jamshid90/todo-list/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type taskHandler struct {
	logger      *zap.Logger
	taskUseCase usecase.TaskService
}

func NewTaskHandler(logger *zap.Logger, taskUseCase usecase.TaskService) chi.Router {
	handler := taskHandler{
		logger:      logger,
		taskUseCase: taskUseCase,
	}

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Post("/", handler.Add())
		r.Get("/{task_id}", handler.Get())
		r.Get("/", handler.List())
	})
	return r
}

// @Router /task [POST]
// @Summary Add
// @Description Task
// @Tags task
// @Accept json
// @Produce json
// @Param body body models.AddTaskRequest true "body"
// @Success 200 {object} models.AddTaskResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *taskHandler) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "task.add"
		reqBody := models.AddTaskRequest{}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			h.logger.Error(op, zap.Error(err))
			render.Render(w, r, models.NewResponseError(400, "invalid_request_body", "Invalid request body"))
			return
		}

		task, err := entity.NewTask(
			uuid.NewString(),
			reqBody.Title,
			reqBody.Priority,
			reqBody.Status,
			reqBody.Description,
			time.Now(),
			time.Now(),
		)

		if err != nil {
			h.logger.Error(op, zap.Error(err))

			switch {
			case errors.Is(err, entity.ErrIncorrectTaskStatus):
				render.Render(w, r, models.NewResponseError(400, "task_status_incorrect", "Invalid task status"))
				return
			case errors.Is(err, entity.ErrIncorrectTaskPriority):
				render.Render(w, r, models.NewResponseError(400, "task_priority_incorrect", "Invalid task Priority"))
				return
			case errors.Is(err, entity.ErrIncorrectTaskPriority):
				render.Render(w, r, models.NewResponseError(400, "task_priority_incorrect", "Invalid task Priority"))
				return
			case errors.Is(err, entity.ErrEmptyTaskTitle):
				render.Render(w, r, models.NewResponseError(400, "task_title_empty", "Task title is empty"))
				return
			default:
				h.logger.Error(op, zap.Error(err))
				render.Render(w, r, models.NewResponseError(500, "internal_error", "Internal error"))
				return
			}
		}
		err = h.taskUseCase.Add(r.Context(), task)
		if err != nil {
			h.logger.Error(op, zap.Error(err))
			render.Render(w, r, models.NewResponseError(400, "invalid_request_body", err.Error()))
			return
		}

		render.JSON(w, r, models.AddTaskResponse{
			Task: models.Task{
				UUID:        task.UUID(),
				Title:       task.Title(),
				Status:      task.Status(),
				Priority:    task.Priority(),
				Description: task.Description(),
				CreatedAt:   task.CreatedAt().String(),
				UpdatedAt:   task.UpdatedAt().String(),
			},
		})
	}
}

// @Router /task/{task_id} [GET]
// @Summary Get
// @Description Get task by task_id
// @Tags task
// @Accept json
// @Produce json
// @Param task_id path string true "task_id"
// @Success 200 {object} models.GetTaskResponse
// @Failure 404 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *taskHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "task.Get"
		taskUUID := chi.URLParam(r, "task_id")
		ctx := r.Context()

		task, err := h.taskUseCase.Task(ctx, taskUUID)
		if err != nil {
			var errNotFound = &entity.ErrNotFound{}
			if errors.As(err, &errNotFound) {
				render.Render(w, r, models.NewResponseError(404, "task_not_found", err.Error()))
				return
			}
			h.logger.Error(op, zap.Error(err), zap.String("task_id", taskUUID))
			render.Render(w, r, models.NewResponseError(500, "internal_error", "Internal error"))
			return
		}

		render.JSON(w, r, models.GetTaskResponse{
			Task: models.Task{
				UUID:        task.UUID(),
				Title:       task.Title(),
				Status:      task.Status(),
				Priority:    task.Priority(),
				Description: task.Description(),
				CreatedAt:   task.CreatedAt().String(),
				UpdatedAt:   task.UpdatedAt().String(),
			},
		})
	}
}

// @Router /task [GET]
// @Summary List
// @Description Task list
// @Tags task
// @Accept json
// @Produce json
// @Param limit query uint64 false "limit" default(10)
// @Param offset query uint64 false "offset" default(0)
// @Param status query string false "status" default(new) Enums(new, in_progress, done)
// @Param priority query string false "priority"  default(high) Enums(urgent, high, normal, low)
// @Success 200 {object} models.ListTaskResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *taskHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "task.List"
		params := models.NewQueryParameter(r.URL.Query())
		list, err := h.taskUseCase.List(r.Context(), params.GetLimit(), params.GetOffset(), params.GetParameters())
		if err != nil {
			h.logger.Error(op, zap.Error(err))
			render.Render(w, r, models.NewResponseError(500, "internal_error", "Internal error"))
			return
		}

		var listTaskResponse = models.ListTaskResponse{}
		for _, item := range list {
			listTaskResponse.List = append(listTaskResponse.List, models.Task{
				UUID:        item.UUID(),
				Title:       item.Title(),
				Status:      item.Status(),
				Priority:    item.Priority(),
				Description: item.Description(),
				CreatedAt:   item.CreatedAt().String(),
				UpdatedAt:   item.UpdatedAt().String(),
			})
		}

		render.JSON(w, r, listTaskResponse)

	}
}

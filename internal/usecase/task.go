package usecase

import (
	"context"
	"errors"

	"github.com/Jamshid90/todo-list/internal/entity"
)

type TaskWrite interface {
	Create(ctx context.Context, task *entity.Task) error
}

type TaskRead interface {
	Task(ctx context.Context, uuid string) (*entity.Task, error)
	All(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Task, error)
}

type TaskService interface {
	Add(ctx context.Context, task *entity.Task) error
	Task(ctx context.Context, uuid string) (*entity.Task, error)
	List(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Task, error)
}

type task struct {
	write TaskWrite
	read  TaskRead
}

func NewTask(write TaskWrite, read TaskRead) (*task, error) {
	if write == nil {
		return nil, errors.New("task writer must not be null")
	}
	if read == nil {
		return nil, errors.New("task writer must not be null")
	}
	return &task{
		write: write,
		read:  read,
	}, nil
}

func (t *task) Add(ctx context.Context, task *entity.Task) error {
	return t.write.Create(ctx, task)
}

func (t *task) Task(ctx context.Context, uuid string) (*entity.Task, error) {
	return t.read.Task(ctx, uuid)
}

func (t *task) List(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Task, error) {
	return t.read.All(ctx, limit, offset, filter)
}

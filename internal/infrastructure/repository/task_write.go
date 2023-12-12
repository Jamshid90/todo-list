package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Jamshid90/todo-list/internal/entity"
	"github.com/Jamshid90/todo-list/internal/pkg/postgres"
	"github.com/jackc/pgconn"
)

type taskWrite struct {
	db *postgres.PostgresDB
}

func NewTaskWrite(db *postgres.PostgresDB) (*taskWrite, error) {
	if db == nil {
		return nil, errors.New("db must not be null")
	}
	return &taskWrite{db}, nil
}

func (t *taskWrite) Create(ctx context.Context, task *entity.Task) error {
	const op = "repository.taskWrite.Create"
	clauses := map[string]interface{}{
		"uuid":        task.UUID(),
		"title":       task.Title(),
		"priority":    task.Priority(),
		"status":      task.Status(),
		"description": task.Description(),
		"created_at":  task.CreatedAt(),
		"updated_at":  task.UpdatedAt(),
	}

	sqlStr, args, err := t.db.Sq.Builder.Insert("task").SetMap(clauses).ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = t.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				return fmt.Errorf("%s: %w", op, entity.NewErrConflict("task"))
			}
			return fmt.Errorf("%s: %w", op, err)
		}
	}
	return nil
}

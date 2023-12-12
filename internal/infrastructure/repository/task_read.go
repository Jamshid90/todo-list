package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Jamshid90/todo-list/internal/entity"
	"github.com/Jamshid90/todo-list/internal/pkg/postgres"
	"github.com/jackc/pgx/v4"
)

type taskRead struct {
	db *postgres.PostgresDB
}

func NewTaskRead(db *postgres.PostgresDB) (*taskRead, error) {
	if db != nil {
		return nil, errors.New("db must not be null")
	}
	return &taskRead{db}, nil
}

func (t *taskRead) Task(ctx context.Context, uuid string) (*entity.Task, error) {
	const op = "repository.taskRead.Task"
	query := t.db.Sq.Builder.
		Select(
			"title",
			"priority",
			"status",
			"description",
			"created_at",
			"updated_at",
		).
		From("task").
		Where(t.db.Sq.Equal("uuid", uuid))

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var (
		title       string
		priority    string
		status      string
		description string
		createdAt   time.Time
		updatedAt   time.Time
	)

	err = t.db.QueryRow(ctx, sqlStr, args...).Scan(
		&title,
		&priority,
		&status,
		&description,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.NewErrNotFound("task", uuid)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	task, err := entity.NewTask(
		uuid,
		title,
		priority,
		status,
		description,
		createdAt,
		updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return task, nil
}

func (t *taskRead) All(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Task, error) {
	const op = "repository.taskRead.All"
	if limit <= 0 {
		limit = 10
	}

	query := t.db.Sq.Builder.
		Select(
			"uuid",
			"title",
			"priority",
			"status",
			"description",
			"created_at",
			"updated_at",
		).
		From("task").
		Limit(limit).
		Offset(offset)

	for k, v := range filter {
		switch k {
		case "status", "priority":
			query = query.Where(t.db.Sq.Equal(k, v))
		}
	}

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := t.db.Query(ctx, sqlStr, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var tasks = make([]*entity.Task, 0, limit)
	for rows.Next() {
		var (
			uuid        string
			title       string
			priority    string
			status      string
			description string
			createdAt   time.Time
			updatedAt   time.Time
		)
		err = rows.Scan(
			&uuid,
			&title,
			&priority,
			&status,
			&description,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		task, err := entity.NewTask(
			uuid,
			title,
			priority,
			status,
			description,
			createdAt,
			updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

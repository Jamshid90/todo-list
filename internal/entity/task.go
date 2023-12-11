package entity

import (
	"errors"
	"time"
)

const (
	TaskStatusNew        = "new"
	TaskStatusIsProgress = "in_progress"
	TaskStatusIsDone     = "done"

	TaskProirityUrgent = "urgent"
	TaskProirityHigh   = "high"
	TaskProirityNormal = "normal"
	TaskProiritLow     = "low"
)

var (
	ErrIncorrectTaskPriority = errors.New("incorrect task priority")
	ErrIncorrectTaskStatus   = errors.New("incorrect task status")

	ErrEmptyTaskUUID       = errors.New("empty task uuid")
	ErrEmptyTaskTitle      = errors.New("empty task title")
	ErrZeroTaskCreatedTime = errors.New("zero task created time")
	ErrZeroTaskUpdatedTime = errors.New("zero task updated time")
)

type Task struct {
	uuid        string
	title       string
	priority    string
	status      string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTask(
	uuid,
	title,
	priority,
	status,
	description string,
	createdAt time.Time,
	updatedAt time.Time,
) (*Task, error) {
	if uuid == "" {
		return nil, ErrEmptyTaskUUID
	}
	if title == "" {
		return nil, ErrEmptyTaskTitle
	}
	if createdAt.IsZero() {
		return nil, ErrZeroTaskCreatedTime
	}
	if updatedAt.IsZero() {
		return nil, ErrZeroTaskUpdatedTime
	}
	task := &Task{
		uuid:        uuid,
		title:       title,
		priority:    priority,
		status:      status,
		description: description,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
	if err := task.IsStatusCorrect(); err != nil {
		return nil, err
	}
	if err := task.IsPriorityCorrect(); err != nil {
		return nil, err
	}
	return task, nil
}

func (t Task) UUID() string {
	return t.uuid
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Status() string {
	return t.status
}

func (t Task) Priority() string {
	return t.priority
}

func (t Task) Description() string {
	return t.description
}

func (t Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t Task) UpdatedAt() time.Time {
	return t.updatedAt
}

// Check task status, if correct return nil
//
// if the status is wrong, return an ErrIncorrectTaskStatus
func (t *Task) IsStatusCorrect() error {
	switch t.status {
	case TaskStatusNew:
		return nil
	case TaskStatusIsProgress:
		return nil
	case TaskStatusIsDone:
		return nil
	}
	return ErrIncorrectTaskStatus
}

// Check task priority, if correct return nil
//
// if the priority is wrong, return an error
func (t *Task) IsPriorityCorrect() error {
	switch t.priority {
	case TaskProirityUrgent:
		return nil
	case TaskProirityHigh:
		return nil
	case TaskProirityNormal:
		return nil
	case TaskProiritLow:
		return nil
	}
	return ErrIncorrectTaskPriority
}

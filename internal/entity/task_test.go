package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	type Task struct {
		Uuid        string
		Title       string
		Priority    string
		Status      string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	type TestCase struct {
		task Task
		name string
		want error
	}

	testCases := []TestCase{
		{
			name: "success",
			task: Task{
				Uuid:        uuid.NewString(),
				Title:       "Add endpint to create task",
				Priority:    "urgent",
				Status:      "new",
				Description: "Some description",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			want: nil,
		},
		{
			name: "empty uuid",
			task: Task{
				Uuid:        "",
				Title:       "Add endpint to create task",
				Priority:    "urgent",
				Status:      "new",
				Description: "Some description",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			want: ErrEmptyTaskUUID,
		},
		{
			name: "empty title",
			task: Task{
				Uuid:        uuid.NewString(),
				Title:       "",
				Priority:    "urgent",
				Status:      "new",
				Description: "Some description",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			want: ErrEmptyTaskTitle,
		},
		{
			name: "incorrect task priority",
			task: Task{
				Uuid:        uuid.NewString(),
				Title:       "Add endpint to create task",
				Priority:    "",
				Status:      "new",
				Description: "Some description",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			want: ErrIncorrectTaskPriority,
		},
		{
			name: "incorrect task status",
			task: Task{
				Uuid:        uuid.NewString(),
				Title:       "Add endpint to create task",
				Priority:    "low",
				Status:      "",
				Description: "Some description",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			want: ErrIncorrectTaskStatus,
		},
		{
			name: "zero task created time",
			task: Task{
				Uuid:     uuid.NewString(),
				Title:    "Add endpint to create task",
				Priority: "low",
				Status:   "new",
			},
			want: ErrZeroTaskCreatedTime,
		},
		{
			name: "zero task updated time",
			task: Task{
				Uuid:      uuid.NewString(),
				Title:     "Add endpint to create task",
				Priority:  "low",
				Status:    "new",
				CreatedAt: time.Now(),
			},
			want: ErrZeroTaskUpdatedTime,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := NewTask(
				testCase.task.Uuid,
				testCase.task.Title,
				testCase.task.Priority,
				testCase.task.Status,
				testCase.task.Description,
				testCase.task.CreatedAt,
				testCase.task.UpdatedAt,
			)
			assert.Equal(t, testCase.want, err)
		})
	}
}

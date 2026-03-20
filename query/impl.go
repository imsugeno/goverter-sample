package query

import (
	"fmt"
	"time"
)

// InMemoryTaskQuery はデモ用のインメモリ実装。
type InMemoryTaskQuery struct{}

var tasks = []TaskResult{
	{
		ID:          "task-001",
		Title:       "OpenAPI仕様を書く",
		Description: "タスク管理APIのOpenAPI仕様を定義する",
		Status:      "done",
		CreatedAt:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:          "task-002",
		Title:       "クエリ層を実装する",
		Description: "CQRSのクエリ側を実装する",
		Status:      "in_progress",
		CreatedAt:   time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:          "task-003",
		Title:       "goverterを導入する",
		Description: "型変換コードを自動生成する",
		Status:      "todo",
		CreatedAt:   time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
	},
}

func (q *InMemoryTaskQuery) ListTasks(input ListTasksInput) ([]TaskResult, error) {
	if input.Status == nil {
		return tasks, nil
	}
	var filtered []TaskResult
	for _, t := range tasks {
		if t.Status == *input.Status {
			filtered = append(filtered, t)
		}
	}
	return filtered, nil
}

func (q *InMemoryTaskQuery) GetTask(input GetTaskInput) (*TaskResult, error) {
	for _, t := range tasks {
		if t.ID == input.TaskID {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("task not found: %s", input.TaskID)
}

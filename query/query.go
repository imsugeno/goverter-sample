package query

// ListTasksInput はタスク一覧取得の入力。
type ListTasksInput struct {
	Status *string
}

// GetTaskInput はタスク詳細取得の入力。
type GetTaskInput struct {
	TaskID string
}

// TaskQuery はタスク参照系のインターフェース。
type TaskQuery interface {
	ListTasks(input ListTasksInput) ([]TaskResult, error)
	GetTask(input GetTaskInput) (*TaskResult, error)
}

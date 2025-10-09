package job

import "time"

type TaskFunc func(params ...any)

type Task struct {
	RunAt  time.Time
	Run    TaskFunc
	Params []any
}
type TaskOption func(*Task)

func NewTask(opts ...TaskOption) *Task {
	task := &Task{}
	for _, opt := range opts {
		opt(task)
	}
	return task
}

func WithRunAt(t time.Time) TaskOption {
	return func(task *Task) {
		task.RunAt = t
	}
}

func WithRun(f TaskFunc) TaskOption {
	return func(task *Task) {
		task.Run = f
	}
}

func WithParams(params ...any) TaskOption {
	return func(task *Task) {
		task.Params = params
	}
}

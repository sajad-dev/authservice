package job

import (
	"sync"
	"time"
)

type TaskFunc func(params ...any)

type Task struct {
	RunAt time.Time 
	Run  TaskFunc 
	Params []any
}


type Jobs struct {
	Tasks []Task
	Mu    sync.Mutex
}

func (j *Jobs) Add(task Task) {
	j.Mu.Lock()
	defer j.Mu.Unlock()

	j.Tasks = append(j.Tasks, task)

}

func (j *Jobs) Run() {
	for {
		if len(j.Tasks) > 0 {
			j.Mu.Lock()
			for i := 0; i < len(j.Tasks); i++ {
				if time.Now().After(j.Tasks[i].RunAt) {
					j.Tasks[i].Run(j.Tasks[i].Params...)
					j.Tasks = append(j.Tasks[:i], j.Tasks[i+1:]...)
					i--
				}
			}
			j.Mu.Unlock()
		}
		time.Sleep(time.Millisecond * 500)

	}
}

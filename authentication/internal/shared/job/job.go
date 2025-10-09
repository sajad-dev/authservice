package job

import (
	"sync"
	"time"
)

type Jobs struct {
	Tasks []*Task
	Mu    sync.Mutex
}

func NewJobs () *Jobs {
	return &Jobs{}
}

type Worker interface {
	Add(task *Task)
	Run()
}

func (j *Jobs) Add(task *Task) {
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

var _ Worker = &Jobs{}

package workers

import "fmt"

type tWorkerData struct {
	Id   int
	Type string
}

type tWorkerComms struct {
	Tasks    chan string
	Stop     chan bool
	Callback chan tWorkerData
}

type Worker struct {
	Data  tWorkerData
	Comms tWorkerComms
}

func NewWorker(id int, t string, tasks chan string) Worker {
	return Worker{
		Data: tWorkerData{
			Id:   id,
			Type: t,
		},
		Comms: tWorkerComms{
			Tasks:    tasks,
			Stop:     make(chan bool),
			Callback: make(chan tWorkerData),
		},
	}
}

func (w Worker) Stop() {
	w.Comms.Stop <- true
}

func (w Worker) Run() {
	fmt.Printf("Worker [%s.%d] starting\n", w.Data.Type, w.Data.Id)
	<-w.Comms.Stop
	fmt.Printf("Worker [%s.%d] stopping\n", w.Data.Type, w.Data.Id)
}

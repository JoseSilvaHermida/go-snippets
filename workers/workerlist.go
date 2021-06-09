package workers

type WorkerList struct {
	List map[string]map[int]Worker
}

func NewWorkerList() WorkerList {
	return WorkerList{
		List: make(map[string]map[int]Worker),
	}
}

func (l *WorkerList) Add(worker Worker) {
	l.List[worker.Data.Type][worker.Data.Id] = worker
}

func (l *WorkerList) Delete(w Worker) {
	delete(l.List[w.Data.Type], w.Data.Id)
}

func (l *WorkerList) Get(t string, id int) Worker {
	return l.List[t][id]
}

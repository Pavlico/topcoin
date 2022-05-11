package worker

type Worker struct {
	RespChan chan string
	Quit     chan bool
}

func NewWorker(maxRoutines int) Worker {
	return Worker{
		RespChan: make(chan string, maxRoutines),
		Quit:     make(chan bool),
	}
}

func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

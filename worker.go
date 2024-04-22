package workerpool

type Worker interface {
	Run()
	Do(func())
	Stop()
}

type worker struct {
	task chan func()
	pool *Pool
}

func NewWorker(p *Pool) Worker {
	return &worker{
		task: make(chan func()),
		pool: p,
	}
}

func (w *worker) run() {
	for f := range w.task {
		f()
		if w.pool != nil {
			w.pool.workers <- w
		}
	}
}

func (w *worker) Run() {
	go w.run()
}

func (w *worker) Do(f func()) {
	w.task <- f
}

func (w *worker) Stop() {
	close(w.task)
}

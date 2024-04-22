package workerpool

type Pool struct {
	workers chan Worker
}

func NewPool(n int) *Pool {
	p := &Pool{
		workers: make(chan Worker, n),
	}
	for range n {
		w := NewWorker(p)
		w.Run()
		p.workers <- w
	}
	return p
}

func (p *Pool) Do(f func()) {
	worker := <-p.workers
	worker.Do(f)
}

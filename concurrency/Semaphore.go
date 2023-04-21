package concurrency

type Semaphore struct {
	sema chan struct{}
}

func (this *Semaphore) acquire() {
	this.acquireN(1)
}

func (this *Semaphore) acquireN(n int) {
	for i := 0; i < n; i++ {
		this.sema <- struct{}{}
	}
}

func (this *Semaphore) release() {
	<-this.sema
}

func CreateSemaphore(size int) *Semaphore {
	s := make(chan struct{}, size)
	return &Semaphore{sema: s}
}

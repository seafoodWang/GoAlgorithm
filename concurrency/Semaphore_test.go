package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	semaphore := CreateSemaphore(2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i+1, semaphore, &wg)
	}
	wg.Wait()
}

func worker(d int, semaphore *Semaphore, wg *sync.WaitGroup) {
	loop := 0
	for {
		fmt.Println(d, "is waiting")
		semaphore.acquire()
		fmt.Println(d, "doing")
		time.Sleep(time.Second * 1)
		semaphore.release()
		fmt.Println(d, "done")
		loop++
		if loop == 100 {
			break
		}
	}
	wg.Done()
}

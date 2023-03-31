package main

import (
	"fmt"
	"sync"
	"time"
)

type MyType struct {
	c        chan int
	isClosed bool
	lock     sync.Mutex
}

func (m *MyType) Add(n int) {
	if m.isOpen() {
		fmt.Println("add data", n)
		m.c <- n
	}
}

func (m *MyType) Close() {
	if m.isOpen() {
		m.lock.Lock()
		defer m.lock.Unlock()
		close(m.c)
		m.isClosed = true
	}
	fmt.Println(m)
}

func (m *MyType) GetData(i int) {
	for {
		data, ok := <-m.c
		fmt.Println(i, "read data", data, ok)
		if !ok {
			return
		}
	}
}

func (m *MyType) isOpen() bool {
	return !m.isClosed
}

func NewMyType(cache int) *MyType {
	tmp := make(chan int, cache)
	return &MyType{
		c:        tmp,
		isClosed: false,
		lock:     sync.Mutex{},
	}
}

func main() {
	//myType := NewMyType(0)
	//fmt.Println(myType)
	//go func(i int) {
	//	myType.GetData(i)
	//}(1)
	//go func(i int) {
	//	myType.GetData(i)
	//}(2)
	//myType.Add(1)
	//myType.Add(2)
	//myType.Add(3)
	//myType.Close()
	//time.Sleep(time.Second)
	//fmt.Println("print myType", myType)

	tmpC := make(chan any, 1000)
	for i := 0; i < 10; i++ {
		tmpC <- struct{}{}
	}
	m := MultiChanClose{
		c:        tmpC,
		isClosed: false,
		Mutex:    sync.Mutex{},
	}

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				m.Add(struct{}{})
			}
			m.Close()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(m.c)
	fmt.Println("hahaha")
}

type MultiChanClose struct {
	c        chan any
	isClosed bool
	sync.Mutex
}

func (m *MultiChanClose) Close() {
	fmt.Println("close")
	if !m.isClosed {
		fmt.Println("close success")
		m.Lock()
		defer m.Unlock()
		close(m.c)
		m.isClosed = true
	}
}

func (m *MultiChanClose) Add(element any) {
	if !m.isClosed {
		m.c <- element
	}
}

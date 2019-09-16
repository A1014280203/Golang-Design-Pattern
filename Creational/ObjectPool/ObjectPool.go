package ObjectPool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// object struct
type Worker struct {
	name int
}

func (w *Worker) run(task string) {
	fmt.Println(w.name, "is working on", task)
	time.Sleep(time.Second * 1)
}

func Employ() *Worker {
	return &Worker{name: rand.Int() % 100}
}

// Pool
type Pool struct {
	sync.Mutex
	inuse     []*Worker
	available []*Worker
	new       func() *Worker
}

func (p *Pool) Bind(n func() *Worker) {
	p.new = n
}

func (p *Pool) Run(task string) {
	w := p.callWorker()
	w.run(task)
	p.letWorkerRest(w)
}

func (p *Pool) callWorker() *Worker {
	p.Lock()
	defer p.Unlock()
	var w *Worker = nil
	if len(p.available) != 0 {
		w = p.available[0]
		p.available = p.available[1:]
	} else {
		w = p.new()
	}
	p.inuse = append(p.inuse, w)
	return w
}

func (p *Pool) letWorkerRest(w *Worker) {
	p.Lock()
	defer p.Unlock()
	p.available = append(p.available, w)
	if w == p.inuse[0] {
		p.inuse = p.inuse[1:]
	} else if w == p.inuse[len(p.inuse)-1] {
		p.inuse = p.inuse[:len(p.inuse)-1]
	} else {
		for i, v := range p.inuse {
			if v == w {
				p.inuse = append(p.inuse[:i], p.inuse[i+1:]...)
				break
			}
		}
	}
}

func Test() {
	p := Pool{}
	p.Bind(Employ)
	go p.Run("task 1")
	p.Run("task 2")
	p.Run("task 3")
}

/*
refer: https://github.com/bvwells/go-patterns/blob/master/creational/object_pool.go
按照go的风格应该将object struct和Employ函数应该单独组织在一个命名空间内
*/

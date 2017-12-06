package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	// pidScanInterval is the interval at which the executor scans the process
	// tree for finding out the pids that the executor and it's child processes
	// have forked
	pidScanInterval = 2 * time.Second
)

type Executor struct {
	pids          map[int]string
	pidLock       sync.RWMutex
	processExited chan interface{}
}

func NewExecutor() *Executor {
	exec := &Executor{
		processExited: make(chan interface{}),
		pids:          make(map[int]string),
	}
	return exec
}

func getAllPids() (map[int]string, error) {
	pids := map[int]string{
		1: "pid1",
		2: "pid2",
	}
	return pids, nil
}

func (e *Executor) collectPids() {
	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("Time:", time.Now())

			pids, err := getAllPids()
			if err != nil {
				fmt.Println(err)
			}
			e.pidLock.Lock()

			for pid, np := range pids {
				if _, ok := e.pids[pid]; !ok {
					e.pids[pid] = np
				}
			}

			for pid := range e.pids {
				if _, ok := pids[pid]; !ok {
					delete(e.pids, pid)
				}
			}

			e.pidLock.Unlock()
			timer.Reset(pidScanInterval)
		case <-e.processExited:
			return
		}
	}
}

func (e *Executor) wait() {
	defer close(e.processExited)

	time.Sleep(10 * time.Second)
}

func (e *Executor) LaunchCmd() error {
	go e.collectPids()
	go e.wait()
	return nil
}

func TestExecutor() {
	exec := NewExecutor()
	err := exec.LaunchCmd()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(20 * time.Second)
}

////////////////////////////////////////////////////////////
func DoTickerWork(res chan interface{}, timeout <-chan time.Time) {
	t := time.NewTicker(3 * time.Second)
	done := make(chan bool, 1)
	go func() {
		defer close(res)
		i := 1
		for {
			select {
			case <-t.C:
				fmt.Printf("start %d th worker\n", i)
				res <- i
				i++
			case <-timeout:
				close(done)
				return
			}
		}
	}()
	<-done
	return
}

func TestTimer() {
	res := make(chan interface{}, 10000)
	timeout := time.After(10 * time.Second)
	DoTickerWork(res, timeout)
	for v := range res {
		fmt.Println(v)
	}
}

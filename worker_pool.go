package main

import "sync"

type Task interface {
	Proccess() string
}

type Result struct {
	TextResult string
}

// Define worker pool
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	taskChan    chan Task
	wg          sync.WaitGroup
	Results     []Result
}

func (wp *WorkerPool) worker() {
	// worker will receive task from channel and execute it
	for task := range wp.taskChan {
		res := task.Proccess()
		wp.Results = append(wp.Results, Result{TextResult: res})

		// send signal to wait group that task already executed
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	// init task channel
	wp.taskChan = make(chan Task, len(wp.Tasks))

	// start worker
	for i := 0; i < wp.concurrency; i++ {
		// execute in go routine
		go wp.worker()
	}

	// put task to channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.taskChan <- task
	}
	close(wp.taskChan)

	// wait until all task finished
	wp.wg.Wait()
}
